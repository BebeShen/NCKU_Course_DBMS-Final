const submitButton = document.getElementById("submit-button")
const sqlQuery = document.getElementById("sql-query")
const sqlInput = document.getElementById("sql-input-command")
const askButton = document.getElementById("ask-button")
const askTime = document.getElementById("ask-time")
const testButton = document.getElementById("test-button")
const searchSelect = document.getElementById("search-select")

submitButton.addEventListener("click", function () {
    console.log("submitButton Click")
    if(sqlQuery.value != ""){
        // use sql query builder
        sqlInput.textContent = sqlQuery.value
    }
    else {
        // use Button
        sqlInput.textContent = searchSelect.value
    }
})

askButton.addEventListener("click", function () {
    console.log("askButton Click")
    askTime.textContent = "new button clicked"
    $.ajax({
        url: "http://localhost:8789/hello", 
        type: "get",
        data: {},
        success: function(result){
            console.log(result);
            $("#div1").html(result);
        }
    });
})

testButton.addEventListener("click", function(){
    console.log("testButton Click")
    var cid = sqlQuery.value
    $.ajax({
        url: "http://localhost:8789/searchCustomerById", 
        type: "post",
        data: {cid},
        success: function(result){
            console.log(result);
            $("#div1").html(result);
        }
    });
});
// $(document).ready(function(){
//     // 找到id為test-button的元件 (#->id)
//     $("#test-button").click(function(){
//         console.log("test");
//         $.ajax({url: "http://localhost:8789/hello", success: function(result){
//             console.log(result);
//             $("#div1").html(result);
//         }});
//     });
// });