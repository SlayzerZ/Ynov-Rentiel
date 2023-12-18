let tasklist = document.getElementById('tasklist')
let add = document.getElementById('newtask')
let checkTask = document.getElementsByTagName('input')
function New(){
    let task = document.createElement('div')
    task.className = "task"
    let inputValue = document.getElementById('taskvalue').value
    task.innerHTML = `<p>${inputValue}</p>` + `<input type="checkbox">`
    if (inputValue == ""){
        alert("Error : You must write something")
    } else{
        tasklist.appendChild(task)
    }
    document.getElementById('taskvalue').value = ""
}
function checkList(){
    let taskall = document.getElementsByClassName('task')
    for (let i = 1; i < checkTask.length; i++){
        if (checkTask[i].checked == true){
            taskall[i-1].remove()
        }
    }
    window.requestAnimationFrame(checkList)
}
add.addEventListener('click', function(){
    New()
    window.requestAnimationFrame(checkList)
})
