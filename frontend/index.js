const submitButton = document.getElementById("submit-button")
const sqlQueryText = document.getElementById("sql-query")
const sqlInput = document.getElementById("sql-input-command")
// const askButton = document.getElementById("ask-button")
// const askTime = document.getElementById("ask-time")
// const testButton = document.getElementById("test-button")
const searchSelect = document.getElementById("search-select")


// ref: https://stackoverflow.com/questions/5180382/convert-json-data-to-a-html-table
function buildHtmlTable(jsonData, selector) {
    var columns = addAllColumnHeaders(jsonData, selector);

    for (var i = 0; i < jsonData.length; i++) {
        var row$ = $('<tr/>');
        for (var colIndex = 0; colIndex < columns.length; colIndex++) {
        var cellValue = jsonData[i][columns[colIndex]];
        if (cellValue == null) cellValue = "";
        row$.append($('<td/>').html(cellValue));
        }
        $(selector).append(row$);
    }
}
  
// Adds a header row to the table and returns the set of columns.
// Need to do union of keys from all records as some records may not contain all records.
function addAllColumnHeaders(jsonData, selector) {
    var columnSet = [];
    var headerTr$ = $('<tr/>');

    for (var i = 0; i < jsonData.length; i++) {
        var rowHash = jsonData[i];
        for (var key in rowHash) {
        if ($.inArray(key, columnSet) == -1) {
            columnSet.push(key);
            headerTr$.append($('<th/>').html(key));
        }
        }
    }
    $(selector).append(headerTr$);

    return columnSet;
}

submitButton.addEventListener("click", function () {
    console.log("submitButton Click")
    if(sqlQueryText.value != ""){
        // use sql query builder "SELECT * FROM customer"
        var sql_query = sqlQueryText.value
        sqlInput.textContent = sqlQueryText.value
        $.ajax({
            url: "http://localhost:8789/queryBuilder", 
            type: "post",
            data: {sql_query},
            success: function(result){
                console.log(result);
                buildHtmlTable(JSON.parse(result), "#data-table")
            }
        });
    }
    else {
        // use Button
        switch (searchSelect.value) {
            case "select-from-where":
                sqlInput.textContent = "SELECT * FROM customer WHERE id=1";
                break;
            default:
                break;
        }
    }
})

// askButton.addEventListener("click", function () {
//     console.log("askButton Click")
//     askTime.textContent = "new button clicked"
//     $.ajax({
//         url: "http://localhost:8789/hello", 
//         type: "get",
//         data: {},
//         success: function(result){
//             console.log(result);
//             $("#div1").html(result);
//         }
//     });
// })

// testButton.addEventListener("click", function(){
//     console.log("testButton Click")
//     var cid = sqlQuery.value
//     $.ajax({
//         url: "http://localhost:8789/searchCustomerById", 
//         type: "post",
//         data: {cid},
//         success: function(result){
//             console.log(result);
//             $("#div1").html(result);
//         }
//     });
// });
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