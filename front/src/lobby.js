let leftBtn = document.querySelector('#leftBtn');
let rightBtn = document.querySelector('#rightBtn');
let rows = document.querySelectorAll('#row');

let ws = new WebSocket('ws://0.0.0.0:8000/lobby');

let rooms = [];
let page  = 0;

ws.onmessage = e => {
    rooms = [];
    const data = JSON.parse(e.data);
    data.rooms.forEach(room => {
        rooms.push(room);
    });
}

function renderRoom(page) {
    const start = rows.children.length * 2 * (page-1) + 1;
    const toShow = rooms.slice(start,start+rows.children.length * 2);
    
    for (row of rows) {
        const twoRooms = toShow.splice(0, 2);
        let htmlInsert = '';
        for (room of twoRooms) {
            htmlInsert += `<td>${room}</td>`
        }
        row.innerHTML = htmlInsert;
    }
}

leftBtn.addEventListener('click', e => {
    if (page>0) {
        page--;
    }
    renderRoom(page);
})

rightBtn.addEventListener('click', e => {
    if (rooms.length > rows.children.length * 2 * page) {
        page++;
    }
    renderRoom(page);
})