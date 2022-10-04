let arr = JSON.parse(localStorage.getItem("array1"))

if (arr !== null) {
    console.log(arr);
    for (let i = 0; i < arr.length; i++) {
        const li = document.createElement("li")
        const task = document.createTextNode(arr[i])
        li.appendChild(task)
        document.querySelector('#myUl').appendChild(li)
        document.querySelector("#myInput").value = ''
        const span = document.createElement("SPAN")
        const txt = document.createTextNode("\u00D7")
        span.className = "close"
        span.appendChild(txt)
        li.append(span)
    }    
} 

function newElement(){
    const li = document.createElement("li")
    const inputValue = document.querySelector('#myInput').value
    const task = document.createTextNode(inputValue)
    arr.push(inputValue)
    li.appendChild(task)
    if(inputValue === ""){
        alert("You must enter something to the field!")
    }else{
        document.querySelector('#myUl').appendChild(li)
    }
    document.querySelector("#myInput").value = ''

    const span = document.createElement("SPAN")
    const txt = document.createTextNode("\u00D7")
    span.className = "close"
    span.appendChild(txt)
    li.append(span)
    localStorage.setItem("array1", JSON.stringify(arr))
}

const list = document.querySelector("ul")
list.addEventListener("click", function (ev){
    ev.target.classList.toggle("checked")
})

const close = document.getElementsByClassName("close")
for(let i = 0; i < close.length; i++){
    close[i].onclick = function(){
        arr.splice(i,1);
        const div = this.parentElement
        div.style.display = "none"
        localStorage.setItem("array1", JSON.stringify(arr))
    }
}