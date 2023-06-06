//@ts-ignore

const form: HTMLFormElement = document.getElementById("form") as HTMLFormElement;
//@ts-ignore
const username: HTMLInputElement = document.getElementById("username") as HTMLInputElement;
//@ts-ignore

const button = document.querySelector("#submit");

if (localStorage.getItem("username")) {
    username.value = localStorage.getItem("username")!
}
//@ts-ignore

const redirect = (e: Event) => {
    e.preventDefault();

    const data: FormData = new FormData(form);
    
    if (data.get("channel") === "" || data.get("username") === "" ) {
        document.getElementById("alert")!.classList.remove("hidden")
        document.getElementById("alert")!.classList.add("block") 
        return;
    } else {
        document.getElementById("alert")!.classList.remove("block")
        document.getElementById("alert")!.classList.add("hidden")
    };

    localStorage.setItem("username", data.get("username") as string);
    //@ts-ignore
    location.replace("/channel?channel="+(data.get("channel")|"public")+"&username="+data.get("username"))
}

form.addEventListener("submit", redirect)