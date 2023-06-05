///ws/{name:[\w\W]}
var queryString = window.location.search;
var queries = (new URLSearchParams(queryString.slice(1)));
var ws = new WebSocket(window.location.href.includes("https") ? "wss" : "ws" + "://" + window.location.host + "/ws?username=" + queries.get("username") + "&channel=" + queries.get("channel"));
//@ts-ignore
var form = document.getElementById("form");
document.getElementById("name").textContent = window.location.pathname.split("/")[2];
var sendMsg = function (e) {
    e.preventDefault();
    var data = new FormData(form);
    if (!localStorage.getItem("username")) {
        location.replace("http://" + window.location.host + "/channel/");
    }
    ws.send(JSON.stringify({
        author: localStorage.getItem("username"),
        content: data.get("msg")
    }));
    form.value = "";
};
form.addEventListener("submit", sendMsg);
ws.onmessage = function (event) {
    var msg = JSON.parse(event.data);
    document.getElementById("messages").innerHTML += "\n      <div class=\"bg-white rounded p-2 mb-2\">\n          <span>\n              ".concat(msg.author, "\n          </span>\n          <p class=\"text-sm\">\n              ").concat(msg.content, "\n          </p>\n      </div>\n    ");
};
console.log(" you are not going to attack this chat");
setInterval(function () { debugger; }, 1000);
