//@ts-ignore
var form = document.getElementById("form");
//@ts-ignore
var username = document.getElementById("username");
//@ts-ignore
var button = document.querySelector("#submit");
if (localStorage.getItem("username")) {
    username.value = localStorage.getItem("username");
}
//@ts-ignore
var redirect = function (e) {
    e.preventDefault();
    var data = new FormData(form);
    if (data.get("channel") === "" || data.get("username") === "") {
        document.getElementById("alert").classList.remove("hidden");
        document.getElementById("alert").classList.add("block");
        return;
    }
    else {
        document.getElementById("alert").classList.remove("block");
        document.getElementById("alert").classList.add("hidden");
    }
    ;
    localStorage.setItem("username", data.get("username"));
    //@ts-ignore
    location.replace("/channel?channel=" + (data.get("channel") | "public") + "&username=" + data.get("username"));
};
form.addEventListener("submit", redirect);
