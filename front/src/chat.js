const chatLog = document.querySelector('.chatLog');
const chatInput = document.querySelector('#chatConsole');
const room = document.URL

ws = new WebSocket(`ws://0.0.0.0:8000/chat/${document.URL.split('/').pop()}`);

ws.onmessage = (e) => {
    
}

$(document).ready(() => {
    // ws = new WebSocket("ws://0.0.0.0:8000/chat/test:1");
    // ws.onmessage = (event) => {
    //     console.log(event);
    //     let data = JSON.parse(event.data);

    //     $("#chatLog").append("<p><span>"+data["msg"]+"</span></p>");
    // }

    lobby = new WebSocket("ws://0.0.0.0:8000/lobby");
    lobby.onmessage = (e) => {
        console.log(e);
    }
});

$("#chatConsole").on("keypress", (e) => {
    if (e.which === 13) {
        let data = $("#chatConsole").val();
        ws.send(data);
        $("#chatConsole").val("");
    }
})
