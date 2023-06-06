/**
 * the json of the message it should to have this structure
 * ```
 * interface {
  Author  string `json:"author"`
  Message string `json:"message"`}
 */
  interface Message {
    author: string;
content: string;
  }
  ///ws/{name:[\w\W]}
  const queryString=window.location.search
  const queries=(new URLSearchParams(queryString.slice(1)))
  const ws = new WebSocket(
    //@ts-ignore
    window.location.href.includes("https")?"wss":"ws"+"://"+window.location.host+ "/ws?username="+queries.get("username")+"&channel="+queries.get("channel")!
  );
  //@ts-ignore
  const form: HTMLFormElement = document.getElementById("form") as HTMLFormElement;
  document.getElementById("name")!.textContent = window.location.pathname.split("/")[2]
  const welcoming="Buenas mis compañeros del MKUltra me acabo de conectar :D "

  ws.onopen=()=>{
    ws.send(JSON.stringify({
        author: queries.get("username"),
        content: welcoming
      }));
  }
  const sendMsg = (e: Event)=> {
    e.preventDefault();
  
    if (!localStorage.getItem("username")) {
      location.replace("http://"+window.location.host+"/channel/")
    }
  
    ws.send(JSON.stringify({
      author: queries.get("username"),
      content: (document.getElementById("message") as HTMLInputElement).value
    }));
    //@ts-ignore
    document.getElementById("messages")!.innerHTML += `
      <div class="bg-white rounded p-2 mb-2">
          <span>
              ${queries.get("username")}
          </span>
          <p class="text-sm">
              ${ (document.getElementById("message") as HTMLInputElement).value}
          </p>
      </div>
    `;
 
 
    (document.getElementById("message") as HTMLInputElement).value="";
  }
  
  form.addEventListener("submit", sendMsg);
  
  ws.onmessage = (event) => {
    let msg: Message = JSON.parse(event.data)
    document.getElementById("messages")!.innerHTML += `
      <div class="bg-white rounded p-2 mb-2">
          <span>
              ${msg.author}
          </span>
          <p class="text-sm">
              ${msg.content}
          </p>
      </div>
    `
  }
  console.log(" you are not going to attack this chat");
  setInterval(()=>{debugger},1000)