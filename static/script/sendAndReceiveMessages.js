///ws/{name:[\w\W]}
var queryString = window.location.search;
var queries = (new URLSearchParams(queryString.slice(1)));
var ws = new WebSocket(
//@ts-ignore
window.location.href.includes("https") ? "wss" : "ws" + "://" + window.location.host + "/ws?username=" + queries.get("username") + "&channel=" + queries.get("channel"));
//@ts-ignore
var form = document.getElementById("form");
document.getElementById("name").textContent = window.location.pathname.split("/")[2];
var welcoming = "Buenas mis compañeros del MKUltra me acabo de conectar :D ";
ws.onopen = function () {
    ws.send(JSON.stringify({
        author: queries.get("username"),
        content: welcoming
    }));
};
var sendMsg = function (e) {
    e.preventDefault();
    if (!localStorage.getItem("username")) {
        location.replace("http://" + window.location.host + "/channel/");
    }
    ws.send(JSON.stringify({
        author: queries.get("username"),
        content: document.getElementById("message").value
    }));
    //@ts-ignore
    document.getElementById("messages").innerHTML += "\n      <div class=\"bg-white rounded p-2 mb-2\">\n          <span>\n              ".concat(queries.get("username"), "\n          </span>\n          <p class=\"text-sm\">\n              ").concat(document.getElementById("message").value, "\n          </p>\n      </div>\n    ");
    document.getElementById("message").value = "";
};
form.addEventListener("submit", sendMsg);
ws.onmessage = function (event) {
    var msg = JSON.parse(event.data);
    document.getElementById("messages").innerHTML += "\n      <div class=\"bg-white rounded p-2 mb-2\">\n          <span>\n              ".concat(msg.author, "\n          </span>\n          <p class=\"text-sm\">\n              ").concat(msg.content, "\n          </p>\n      </div>\n    ");
};
console.log(" you are not going to attack this chat");
setInterval(function () { debugger; }, 1000);
