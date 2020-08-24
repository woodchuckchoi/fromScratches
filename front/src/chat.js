const chatLog   = document.querySelector('.chatTable');
const chatInput = document.querySelector('#chatConsole');
const sendBtn   = document.querySelector('#sendButton');

const room = document.URL;
const nameLength = 20;

document.querySelector('.title').innerText = `Room test1`

let ws = new WebSocket(`ws://0.0.0.0:8000/chat/test1`);
console.log(ws)

ws.onmessage = (e) => {
    const data = JSON.parse(e.data);
    const name = data['sender']+' '.repeat(nameLength-data['sender'].length);
    const message = data['msg'];
    if (chatLog.children.length > 100) {
        chatLog.children[0].remove();
    }
    chatLog.innerHTML += `<td class="logName">${name}</td><td class="logMessage">${message}</td>`;
}

chatInput.addEventListener('keyup', (e) => {
    const key = 'which' in e ? e.which : e.keyCode;
    if (key === 13) {
        sendData();
    }
})

sendBtn.addEventListener('click', (e) => {
    sendData();
})

function sendData() {
    let data = chatInput.value;
    ws.send(data);
    chatInput.value = '';
}