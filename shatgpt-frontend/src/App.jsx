import { useState } from 'react'
import reactLogo from './assets/react.svg'
import './App.css'

import Leaderboard from './components/leaderboard'
import axios from 'axios'


function App() {
    const [count, setCount] = useState(0)

    function promptToGTP() {
        let inputToGTP = prompt("What would you ask ChatGTP?")
        let charCount = inputToGTP.length
        if (count > charCount) {
            return;
        }
        else {
            alert("You do not have enough coins!")
        }
    }

    window.onload = function () {
        let conn;
        let msg = document.getElementById("msg");
        let log = document.getElementById("log");

        function appendLog(item) {
            let doScroll = log.scrollTop > log.scrollHeight - log.clientHeight - 1;
            log.appendChild(item);
            if (doScroll) {
                log.scrollTop = log.scrollHeight - log.clientHeight;
                window.onload = function () {
                    if (localStorage.getItem("user-token") === null) {
                        axios.post('/api/register', {
                            desired_name: ''
                        })
                            .then(function (response) {
                                console.log(response.data);
                                localStorage.setItem("user-token", response.data.token)
                            })
                            .catch(function (error) {
                                console.log(error);
                            });
                    }


                    let conn;
                    let msg = document.getElementById("msg");
                    let log = document.getElementById("log");

                    function appendLog(item) {
                        let doScroll = log.scrollTop > log.scrollHeight - log.clientHeight - 1;
                        log.appendChild(item);
                        if (doScroll) {
                            log.scrollTop = log.scrollHeight - log.clientHeight;
                        }
                    }

                    function removeFadeOut(el, speed) {
                        var seconds = speed / 1000;
                        el.style.transition = "all " + seconds + "s ease";
                        el.style.left = "-100vh";
                        el.style.opacity = 0;

                        setTimeout(function () {
                            el.parentNode.removeChild(el);
                        }, speed);
                    }

                    document.getElementById("form").onsubmit = function () {
                        if (!conn) {
                            return false;
                        }

                        // let send_to_socket = {
                        //     event_name: "press",
                        //     token: localStorage.getItem("user-token")
                        // }

                        conn.send(send_to_socket);

                        return false;
                    };

                    if (window["WebSocket"]) {
                        // const params = window.location.href.split("/");
                        // const roomId = params[params.length - 1];
                        var websock_proto = "wss";
                        if (location.protocol !== "https:") {
                            websock_proto = "ws";
                        }
                        conn = new WebSocket(websock_proto + "://" + document.location.host + "/ws/shatgpt");
                        conn.onclose = function (evt) {
                            let item = document.createElement("div");
                            item.innerHTML = "<b>Connection closed.</b>";
                            appendLog(item);
                        };
                        conn.onmessage = function (evt) {
                            let messages = evt.data.split('\n');
                            for (let i = 0; i < messages.length; i++) {
                                let item = document.createElement("div");
                                item.className = "poop";
                                item.innerText = messages[i];
                                appendLog(item);
                                setTimeout(() => {
                                    removeFadeOut(item, 1000)
                                    //item.remove();
                                }, 10000)
                            }
                        };
                    } else {
                        let item = document.createElement("div");
                        item.innerHTML = "<b>Your browser does not support WebSockets.</b>";
                        appendLog(item);
                    }

                };

                if (window["WebSocket"]) {
                    // const params = window.location.href.split("/");
                    // const roomId = params[params.length - 1];
                    var websock_proto = "wss";
                    if (location.protocol !== "https:") {
                        websock_proto = "ws";
                    }
                    conn = new WebSocket(websock_proto + "://" + document.location.host + "/ws/shatgpt");
                    conn.onclose = function (evt) {
                        let item = document.createElement("div");
                        item.innerHTML = "<b>Connection closed.</b>";
                        appendLog(item);
                    };
                    conn.onmessage = function (evt) {
                        let messages = evt.data.split('\n');
                        for (let i = 0; i < messages.length; i++) {
                            let item = document.createElement("div");
                            item.className = "poop";
                            item.innerText = messages[i];
                            appendLog(item);
                            setTimeout(() => {
                                removeFadeOut(item, 1000)
                                //item.remove();
                            }, 10000)
                        }
                    };
                } else {
                    let item = document.createElement("div");
                    item.innerHTML = "<b>Your browser does not support WebSockets.</b>";
                    appendLog(item);
                }

            };

            return (
                <>
                    <div id="log"></div>
                    <form id="form">
                        {/* <input type="text" id="msg" size="64" autofocus/>
      <input type="submit" value="🐘" size="50"/> */}
                        <button type="submit" className="button-36" role="button">
                            💩
                        </button>
                        <button type="submit" className="button-36" role="button" onClick={promptToGTP}>
                            Ask ChatGTP
                        </button>
                    </form>
                    <Leaderboard />
                </>
            );
            return (
                <>
                    <div id="log"></div>
                    <form id="form">
                        {/* <input type="text" id="msg" size="64" autofocus/>
      <input type="submit" value="🐘" size="50"/> */}
                        <button type="submit" class="button-36" role="button">💩</button>
                    </form>
                </>
            )
        }
    }
}

export default App
