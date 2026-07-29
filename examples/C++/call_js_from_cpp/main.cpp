// Call JavaScript from C++ Example

#include "webui.hpp"
#include <iostream>
#include <sstream>
#include <string>

class CounterApp {
public:
    CounterApp() {
        // Bind HTML element IDs with class methods
        win.bind("my_function_count", this, &CounterApp::my_function_count);
        win.bind("Exit", this, &CounterApp::exit_app);
    }

    void run() {
        const std::string html = R"V0G0N(
        <!DOCTYPE html>
        <html lang="en">
        <head>
            <meta charset="UTF-8">
            <meta name="viewport" content="width=device-width, initial-scale=1.0">
            <script src="/webui.js"></script>
            <title>Call JavaScript from C++ Example</title>
            <style>
                :root {
                    --bg-top: #0f172a;
                    --bg-bottom: #020617;
                    --surface: #1e293b;
                    --surface-hover: #334155;
                    --border: #334155;
                    --border-hover: #475569;
                    --accent: #38bdf8;
                    --danger: #ef4444;
                    --danger-hover: #dc2626;
                    --text-main: #f8fafc;
                    --text-muted: #94a3b8;
                }

                * {
                    box-sizing: border-box;
                    margin: 0;
                    padding: 0;
                }

                body {
                    background: linear-gradient(to top, var(--bg-top), var(--bg-bottom));
                    background-attachment: fixed;
                    color: var(--text-main);
                    font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial, sans-serif;
                    min-height: 100vh;
                    padding: 60px 20px;
                    display: flex;
                    flex-direction: column;
                    align-items: center;
                }

                header {
                    width: 100%;
                    max-width: 640px;
                    margin-bottom: 32px;
                }

                h1 {
                    font-size: 1.5rem;
                    font-weight: 600;
                    letter-spacing: -0.025em;
                }

                main {
                    width: 100%;
                    max-width: 640px;
                    display: flex;
                    flex-direction: column;
                    gap: 28px;
                }

                .counter-card {
                    background-color: var(--surface);
                    border: 1px solid var(--border);
                    border-radius: 8px;
                    padding: 32px;
                    text-align: center;
                }

                .counter-label {
                    font-size: 0.75rem;
                    font-weight: 600;
                    text-transform: uppercase;
                    letter-spacing: 0.05em;
                    color: var(--text-muted);
                    margin-bottom: 8px;
                }

                .counter-value {
                    font-size: 4rem;
                    font-weight: 700;
                    color: var(--accent);
                    font-variant-numeric: tabular-nums;
                    line-height: 1;
                }

                .section-title {
                    font-size: 0.75rem;
                    font-weight: 600;
                    text-transform: uppercase;
                    letter-spacing: 0.05em;
                    color: var(--text-muted);
                    margin-bottom: 12px;
                }

                .grid {
                    display: grid;
                    grid-template-columns: repeat(auto-fit, minmax(190px, 1fr));
                    gap: 12px;
                }

                button {
                    background-color: var(--surface);
                    color: var(--text-main);
                    border: 1px solid var(--border);
                    padding: 12px 16px;
                    border-radius: 6px;
                    font-size: 0.875rem;
                    font-weight: 500;
                    font-family: inherit;
                    cursor: pointer;
                    transition: background-color 0.15s ease, border-color 0.15s ease, opacity 0.15s ease;
                    text-align: center;
                }

                button:hover:not(:disabled) {
                    background-color: var(--surface-hover);
                    border-color: var(--border-hover);
                }

                button:active:not(:disabled) {
                    background-color: var(--surface);
                }

                button:disabled {
                    opacity: 0.5;
                    cursor: not-allowed;
                }

                button.primary {
                    background-color: #2563eb;
                    border-color: #2563eb;
                }

                button.primary:hover:not(:disabled) {
                    background-color: #1d4ed8;
                    border-color: #1d4ed8;
                }

                button.danger {
                    background-color: transparent;
                    border-color: var(--danger);
                    color: var(--danger);
                }

                button.danger:hover:not(:disabled) {
                    background-color: var(--danger);
                    color: white;
                }
            </style>
        </head>
        <body>
            <header>
                <h1>WebUI - Call JavaScript from C++</h1>
            </header>

            <main>
                <div class="counter-card">
                    <div class="counter-label">Current Count</div>
                    <div class="counter-value" id="count">0</div>
                </div>

                <section>
                    <div class="section-title">Controls</div>
                    <div class="grid">
                        <button id="ManualBtn" class="primary" onclick="my_function_count();">Manual Count</button>
                        <button id="AutoBtn" onclick="AutoTest();">Auto Count (Every 10ms)</button>
                        <button id="Exit" class="danger" onclick="this.disabled=true;">Exit</button>
                    </div>
                </section>
            </main>

            <script>
                let count = 0;
                let auto_running = false;

                function GetCount() {
                    return count;
                }

                function SetCount(number) {
                    document.getElementById('count').innerHTML = number;
                    count = number;
                }

                function AutoTest(number) {
                    if (auto_running) return;
                    auto_running = true;
                    document.getElementById('AutoBtn').disabled = true;
                    document.getElementById('ManualBtn').disabled = true;

                    setInterval(function() {
                        my_function_count();
                    }, 10);
                }
            </script>
        </body>
        </html>
        )V0G0N";

        // Set WebUI configuration to process UI events one at a time.
        // This example calls C++ from JavaScript every 10ms, so it's 
        // recommended to set this to `true` to avoid race conditions.
        webui::set_config(ui_event_blocking, true);

        // Show the HTML page in the WebView/Browser window
        win.show(html);
            // win.show_browser(html, webui_browser::Chrome);
            // win.show_wv(html);

        // Wait until the window gets closed
        webui::wait();
    }

private:
    webui::window win;

    void my_function_count(webui::window::event* e) {
        // Create a buffer to hold the response
        char response[64];

        // Run JavaScript to get the current count value from the HTML page
        if (!e->get_window().script("return GetCount();", 0, response, 64)) {
            if (!e->get_window().is_shown())
                std::cout << "Window closed." << std::endl;
            else
                std::cout << "JavaScript Error: " << response << std::endl;
            return;
        }

        // Get the count, increment, and send it back
        int count = std::stoi(response) + 1;
        std::stringstream js;
        js << "SetCount(" << count << ");";
        // Run JavaScript (no response) to update the count in the HTML page
        e->get_window().run(js.str());
    }

    void exit_app(webui::window::event* e) {
        // Exit the application, and close all windows. `webui::wait()` will return (Break).
        webui::exit();
    }
};

int main() {
    CounterApp app;
    app.run();
    return 0;
}

#if defined(_MSC_VER)
int APIENTRY WinMain(HINSTANCE hInst, HINSTANCE hInstPrev, PSTR cmdline, int cmdshow) { return main(); }
#endif
