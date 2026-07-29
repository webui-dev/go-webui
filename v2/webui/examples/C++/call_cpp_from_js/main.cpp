// Call C++ from JavaScript Example

#include "webui.hpp"
#include <iostream>

class CallCppFromJsApp {
public:
    CallCppFromJsApp() {
        // Bind HTML element IDs with class methods
        win.bind("my_function_string", this, &CallCppFromJsApp::my_function_string);
        win.bind("my_function_integer", this, &CallCppFromJsApp::my_function_integer);
        win.bind("my_function_boolean", this, &CallCppFromJsApp::my_function_boolean);
        win.bind("my_function_with_response", this, &CallCppFromJsApp::my_function_with_response);
    }

    void run() {
        const std::string html = R"V0G0N(
        <!DOCTYPE html>
        <html lang="en">
        <head>
            <meta charset="UTF-8">
            <meta name="viewport" content="width=device-width, initial-scale=1.0">
            <script src="/webui.js"></script>
            <title>Call C++ from JavaScript Example</title>
            <style>
                :root {
                    --bg-top: #0f172a;
                    --bg-bottom: #020617;
                    --surface: #1e293b;
                    --surface-hover: #334155;
                    --border: #334155;
                    --border-hover: #475569;
                    --accent: #38bdf8;
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
                    margin-bottom: 40px;
                }

                h1 {
                    font-size: 1.5rem;
                    font-weight: 600;
                    letter-spacing: -0.025em;
                    margin-bottom: 6px;
                }

                .subtitle {
                    color: var(--text-muted);
                    font-size: 0.875rem;
                }

                main {
                    width: 100%;
                    max-width: 640px;
                    display: flex;
                    flex-direction: column;
                    gap: 32px;
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
                    transition: background-color 0.15s ease, border-color 0.15s ease;
                    text-align: left;
                    display: flex;
                    align-items: center;
                    justify-content: space-between;
                }

                button:hover {
                    background-color: var(--surface-hover);
                    border-color: var(--border-hover);
                }

                button:active {
                    background-color: var(--surface);
                }

                button.primary {
                    background-color: #2563eb;
                    border-color: #2563eb;
                }

                button.primary:hover {
                    background-color: #1d4ed8;
                    border-color: #1d4ed8;
                }

                .action-row {
                    display: flex;
                    gap: 12px;
                    align-items: center;
                }

                .action-row button {
                    flex: 1;
                }

                .input-wrapper {
                    display: flex;
                    align-items: center;
                    background-color: var(--surface);
                    border: 1px solid var(--border);
                    border-radius: 6px;
                    padding: 0 14px;
                    height: 43px;
                    width: 180px;
                }

                .input-wrapper label {
                    font-size: 0.875rem;
                    color: var(--text-muted);
                    margin-right: 10px;
                    user-select: none;
                }

                input[type="text"] {
                    background: transparent;
                    border: none;
                    color: var(--accent);
                    font-size: 0.875rem;
                    font-weight: 600;
                    font-family: inherit;
                    width: 100%;
                    outline: none;
                }
            </style>
        </head>
        <body>
            <header>
                <h1>WebUI - Call C++ from JavaScript</h1>
                <p class="subtitle">Call C++ functions with arguments <em>(See logs in terminal)</em></p>
            </header>

            <main>
                <section>
                    <div class="section-title">Standard Execution</div>
                    <div class="grid">
                        <button onclick="my_function_string('Hello', 'World');">
                            <span>my_function_string()</span>
                        </button>
                        <button onclick="my_function_integer(123, 456, 789);">
                            <span>my_function_integer()</span>
                        </button>
                        <button onclick="my_function_boolean(true, false);">
                            <span>my_function_boolean()</span>
                        </button>
                    </div>
                </section>

                <section>
                    <div class="section-title">Response Handling</div>
                    <div class="action-row">
                        <button class="primary" onclick="MyJS();">
                            <span>my_function_with_response()</span>
                        </button>
                        <div class="input-wrapper">
                            <label for="MyInputID">Double:</label>
                            <input type="text" id="MyInputID" value="2">
                        </div>
                    </div>
                </section>
            </main>

            <script>
                function MyJS() {
                    const MyInput = document.getElementById('MyInputID');
                    const number = MyInput.value;
                    my_function_with_response(number, 2).then((response) => {
                        MyInput.value = response;
                    });
                }
            </script>
        </body>
        </html>
        )V0G0N";

        // Show the HTML page in the WebView/Browser window
        win.show(html);
            // win.show_browser(html, webui_browser::Chrome);
            // win.show_wv(html);

        // Wait until the window gets closed
        webui::wait();
    }

private:
    webui::window win;

    void my_function_string(webui::window::event* e) {
        // JavaScript: my_function_string('Hello', 'World');
        std::string str_1 = e->get_string(); // Or e->get_string(0);
        std::string str_2 = e->get_string(1);
        std::cout << "my_function_string 1: " << str_1 << std::endl; // Hello
        std::cout << "my_function_string 2: " << str_2 << std::endl; // World
    }

    void my_function_integer(webui::window::event* e) {
        // JavaScript: my_function_integer(123, 456, 789);
        long long number_1 = e->get_int(); // Or e->get_int(0);
        long long number_2 = e->get_int(1);
        long long number_3 = e->get_int(2);
        std::cout << "my_function_integer 1: " << number_1 << std::endl; // 123
        std::cout << "my_function_integer 2: " << number_2 << std::endl; // 456
        std::cout << "my_function_integer 3: " << number_3 << std::endl; // 789
    }

    void my_function_boolean(webui::window::event* e) {
        // JavaScript: my_function_boolean(true, false);
        bool status_1 = e->get_bool(); // Or e->get_bool(0);
        bool status_2 = e->get_bool(1);
        std::cout << "my_function_boolean 1: " << (status_1 ? "True" : "False") << std::endl;
        std::cout << "my_function_boolean 2: " << (status_2 ? "True" : "False") << std::endl;
    }

    void my_function_with_response(webui::window::event* e) {
        // JavaScript: my_function_with_response(number, 2).then(...)
        long long number = e->get_int(0);
        long long times = e->get_int(1);
        long long res = number * times;
        std::cout << "my_function_with_response: " << number << " * " << times << " = " << res << std::endl;
        // Send the response back to JavaScript
        e->return_int(res);
    }
};

int main() {
    CallCppFromJsApp app;
    app.run();
    return 0;
}

#if defined(_MSC_VER)
int APIENTRY WinMain(HINSTANCE hInst, HINSTANCE hInstPrev, PSTR cmdline, int cmdshow) { return main(); }
#endif
