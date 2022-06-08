const submitButton = document.getElementById("submit-button")
const sqlQueryText = document.getElementById("sql-query")
const sqlInput = document.getElementById("sql-input-command")
const searchSelect = document.getElementById("search-select")
const messageText = document.getElementById("data-message")


// ref: https://stackoverflow.com/questions/5180382/convert-json-data-to-a-html-table
function buildHtmlTable(jsonData, selector) {
    // clean table
    document.getElementById("data-table").innerHTML = '';

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

function handleChange() {
    console.log(searchSelect.value)
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
                sqlInput.textContent = "SELECT * FROM customer";
                $.ajax({
                    url: "http://localhost:8789/getAllCustomer", 
                    type: "get",
                    data: {},
                    success: function(result){
                        console.log(result);
                        buildHtmlTable(JSON.parse(result), "#data-table")
                    }
                });
                break;
            case "insert":
                sqlInput.textContent = "INSERT INTO customer (username, password, c_location) VALUES ('test', 'test', '701台南市東區莊敬里 中華東路一段 66號');";
                $.ajax({
                    url: "http://localhost:8789/addCustomer", 
                    type: "post",
                    data: {
                        username:"test",
                        password:"test",
                        c_location:"701台南市東區莊敬里 中華東路一段 66號"
                    },
                    success: function(result){
                        console.log(result);
                        messageText.textContent = result;
                    }
                });
                break;
            case "update":
                sqlInput.textContent = "UPDATE Customer SET c_location='701台南市東區大學路1號' WHERE c_id = 1;";
                $.ajax({
                    url: "http://localhost:8789/updateCustomer", 
                    type: "post",
                    data: {
                        c_id:1,
                        c_location:"701台南市東區大學路1號"
                    },
                    success: function(result){
                        console.log(result);
                        messageText.textContent = result;
                    }
                });
                break;
            case "delete":
                sqlInput.textContent = "DELETE FROM Customer WHERE c_id=1;";
                $.ajax({
                    url: "http://localhost:8789/deleteCustomer", 
                    type: "post",
                    data: {
                        c_id:1
                    },
                    success: function(result){
                        console.log(result);
                        messageText.textContent = result;
                    }
                });
                break;
            default:
                break;
        }
    }
})
