window.onload =()=> {
    let log = document.getElementsByClassName(`log`)[0];
    log.innerHTML = "Welcome<b>";

    let btns = document.getElementsByTagName("button");
    let ConnWorld1 = btns[0];
    let ConnWorld2 = btns[1];
    let SendData = btns[2];
    var conn;
    ConnWorld1.addEventListener(`click`, (e)=>{
        conn = new WebSocket("ws://localhost:8888/test1");
        conn.onmessage =(e)=>{
            log.innerHTML += e.data+`</br>`;
        }
    })
    ConnWorld2.addEventListener(`click`, (e)=>{
        conn = new WebSocket("ws://localhost:8888/test2");
        conn.onmessage =(e)=>{
            log.innerHTML += e.data+`</br>`;
        }
    })
    SendData.addEventListener(`click`, (e)=>{
        conn.send("Hello!, Here's some random data->" +Math.round(Math.random()*10000))
    })
}