const submitButton = document.getElementById("submit-button")
const sqlQueryText = document.getElementById("sql-query")
const sqlInput = document.getElementById("sql-input-command")
const sqldes = document.getElementById("sql-des")
const searchSelect = document.getElementById("search-select")
const searchSelect2 = document.getElementById("search2-select")
const messageText = document.getElementById("data-message")

const selectOptions = {
    "select-from-where":["customer","food","wasted","store","order"], 
    "insert":["customer","employee","food"], 
    "update":["customer"],
    "delete":["customer"], 
    "in":["food"],
    "not-in":["food"],
    "exists":["food"],
    "not-exists":["food"],
    "count":["food"],
    "sum":["food"],
    "max":["food"],
    "min":["food"],
    "avg":["food"],
    "having":["food"],
};

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
    var options_str = "";
    console.log(searchSelect.value)
    searchSelect2.innerHTML = "";
    for (let index = 0; index < selectOptions[searchSelect.value].length; index++) {
        options_str += '<option value="' + selectOptions[searchSelect.value][index] + '">' + selectOptions[searchSelect.value][index] + '</option>';
    }
    searchSelect2.innerHTML = options_str;
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
        sqldes.textContent = ""
        // use Button
        switch (searchSelect.value) {
            case "select-from-where":
                var req_url = "http://localhost:8789/getAllCustomer";
                switch (searchSelect2.value) {
                    case "customer":
                        sqlInput.textContent = "SELECT * FROM Customer";
                        req_url = "http://localhost:8789/getAllCustomer";
                        break;
                    case "employee":
                        sqlInput.textContent = "SELECT * FROM Employee";
                        req_url = "http://localhost:8789/getAllEmployee";
                        break;
                    case "food":
                        sqlInput.textContent = "SELECT * FROM Food";
                        req_url = "http://localhost:8789/getAllFood";
                        break;
                    case "wasted":
                        sqlInput.textContent = "SELECT * FROM Wasted";
                        req_url = "http://localhost:8789/getAllWasted";
                        break;
                    case "store":
                        sqlInput.textContent = "SELECT * FROM Store";
                        req_url = "http://localhost:8789/getAllStore";
                        break;
                    case "order":
                        sqlInput.textContent = `SELECT c.username AS 暱稱, f.name AS 食物名稱, s.name AS 商店名稱, o.message, o.status
                        FROM Orders AS o
                        LEFT JOIN Customer AS c ON o.c_id = c.c_id
                        LEFT JOIN Food AS f on o.f_id = f.f_id
                        LEFT JOIN Store AS s on o.s_id = s.s_id`;
                        req_url = "http://localhost:8789/getAllOrder";
                        break;
                    default:
                        break;
                }
                $.ajax({
                    url: req_url, 
                    type: "get",
                    data: {},
                    success: function(result){
                        console.log(result);
                        buildHtmlTable(JSON.parse(result), "#data-table")
                    }
                });
                break;
            case "insert":
                var req_url = "http://localhost:8789/addCustomer";
                var req_data = {};
                switch (searchSelect2.value) {
                    case "customer":
                        sqlInput.textContent = "INSERT INTO Customer (username, password, c_location) VALUES ('test', 'test', '701台南市東區莊敬里 中華東路一段 66號');";
                        req_url = "http://localhost:8789/addCustomer";
                        req_data = {
                            username:"test",
                            password:"test",
                            c_location:"701台南市東區莊敬里 中華東路一段 66號"
                        };
                        break;
                    case "employee":
                        sqlInput.textContent = "INSERT INTO Employee(username, password) VALUES ('employee', 'employee_password')";
                        req_url = "http://localhost:8789/addEmployee";
                        req_data = {
                            username:"employee",
                            password:"employee_password",
                            work_for:1
                        }
                        break;
                    case "food":
                        sqlInput.textContent = "INSERT INTO Food(category, name, expireDate, price, discount) VALUES ('riceroll', '炙燒明太子鮭魚飯糰', '2022-06-12', 33, 0)";
                        req_url = "http://localhost:8789/addFood";
                        req_data = {
                            category: "riceroll",
                            name: "炙燒明太子鮭魚飯糰",
                            expireDate: "2022-06-12",
                            price: 33,
                            discount: 0,
                            store_at: 1
                        }
                        break;
                    default:
                        break;
                }
                $.ajax({
                    url: req_url, 
                    type: "post",
                    data: req_data,
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
                sqlInput.textContent = "DELETE FROM Customer WHERE c_id=9;";
                $.ajax({
                    url: "http://localhost:8789/deleteCustomer", 
                    type: "post",
                    data: {
                        c_id:9
                    },
                    success: function(result){
                        console.log(result);
                        messageText.textContent = result;
                    }
                });
                break;
            case "in":
                sqldes.textContent = "查詢在商店1中飯糰和麵包類的所有食物"
                sqlInput.textContent = `SELECT * FROM Food WHERE category IN ('riceroll','bread') AND store_at = 1`;
                $.ajax({
                    url: "http://localhost:8789/queryBuilder", 
                    type: "post",
                    data: {
                        sql_query:sqlInput.textContent
                    },
                    success: function(result){
                        console.log(result);
                        buildHtmlTable(JSON.parse(result), "#data-table")
                    }
                });
                break;
            case "not-in":
                sqldes.textContent = "查詢在商店1中尚未被預約的所有食物"
                sqlInput.textContent = `SELECT * FROM Food WHERE f_id NOT IN (SELECT f_id FROM Orders) AND store_at = 1`;
                $.ajax({
                    url: "http://localhost:8789/queryBuilder", 
                    type: "post",
                    data: {
                        sql_query:sqlInput.textContent
                    },
                    success: function(result){
                        console.log(result);
                        buildHtmlTable(JSON.parse(result), "#data-table")
                    }
                });
                break;
            case "exists":
                sqldes.textContent = "查詢在商店1中已經被預約的所有食物"
                sqlInput.textContent = `SELECT * FROM Food WHERE EXISTS (SELECT f_id FROM Orders WHERE Orders.f_id=Food.f_id)`;
                $.ajax({
                    url: "http://localhost:8789/queryBuilder", 
                    type: "post",
                    data: {
                        sql_query:sqlInput.textContent
                    },
                    success: function(result){
                        console.log(result);
                        buildHtmlTable(JSON.parse(result), "#data-table")
                    }
                });
                break;
            case "not-exists":
                sqldes.textContent = "查詢在商店1中尚未預約的所有食物"
                sqlInput.textContent = `SELECT * FROM Food WHERE NOT EXISTS (SELECT f_id FROM Orders WHERE Orders.f_id=Food.f_id )`;
                $.ajax({
                    url: "http://localhost:8789/queryBuilder", 
                    type: "post",
                    data: {
                        sql_query:sqlInput.textContent
                    },
                    success: function(result){
                        console.log(result);
                        buildHtmlTable(JSON.parse(result), "#data-table")
                    }
                });
                break;
            case "count":
                sqldes.textContent = "查詢在商店1的訂單個數"
                sqlInput.textContent = `SELECT COUNT(*) FROM Orders WHERE s_id = 1`;
                $.ajax({
                    url: "http://localhost:8789/queryBuilder", 
                    type: "post",
                    data: {
                        sql_query:sqlInput.textContent
                    },
                    success: function(result){
                        console.log(result);
                        buildHtmlTable(JSON.parse(result), "#data-table")
                    }
                });
                break;
            case "sum":
                sqldes.textContent = "查詢顧客1的所有訂單金額"
                sqlInput.textContent = `
                SELECT SUM(f.price) AS price 
                FROM Orders AS o LEFT JOIN Food AS f ON o.f_id = f.f_id 
                WHERE o.c_id = 1`;
                $.ajax({
                    url: "http://localhost:8789/queryBuilder", 
                    type: "post",
                    data: {
                        sql_query:sqlInput.textContent
                    },
                    success: function(result){
                        console.log(result);
                        buildHtmlTable(JSON.parse(result), "#data-table")
                    }
                });
                break;
            case "max":
                sqldes.textContent = "查詢在所有商店尚未預約食物最多的商店"
                sqlInput.textContent = `
                SELECT name,MAX(left_food) AS num_left_food 
                FROM ( 
                  SELECT Store.name AS name, COUNT(store_at) AS left_food 
                  FROM Food 
                  LEFT JOIN Store ON Food.store_at = Store.s_id 
                  GROUP BY store_at 
                )
                `;
                $.ajax({
                    url: "http://localhost:8789/queryBuilder", 
                    type: "post",
                    data: {
                        sql_query:sqlInput.textContent
                    },
                    success: function(result){
                        console.log(result);
                        buildHtmlTable(JSON.parse(result), "#data-table")
                    }
                });
                break;
            case "min":
                sqldes.textContent = "查詢在商店1中目前最接近保存期限的食物"
                sqlInput.textContent = `SELECT Food.f_id, Food.name, MIN(expireDate) AS expiredDate FROM Food WHERE store_at = 1`;
                $.ajax({
                    url: "http://localhost:8789/queryBuilder", 
                    type: "post",
                    data: {
                        sql_query:sqlInput.textContent
                    },
                    success: function(result){
                        console.log(result);
                        buildHtmlTable(JSON.parse(result), "#data-table")
                    }
                });
                break;
            case "avg":
                sqldes.textContent = "查詢在商店1中飯糰的平均價格"
                sqlInput.textContent = `SELECT category, AVG(price) as price FROM Food WHERE category="riceroll"`;
                $.ajax({
                    url: "http://localhost:8789/queryBuilder", 
                    type: "post",
                    data: {
                        sql_query:sqlInput.textContent
                    },
                    success: function(result){
                        console.log(result);
                        buildHtmlTable(JSON.parse(result), "#data-table")
                    }
                });
                break;
            case "having":
                sqldes.textContent = "查詢在所有商店中剩餘食物個數大於5的商店"
                sqlInput.textContent = `
                SELECT s.name, COUNT(f_id) AS num_food_left
                FROM Food AS f
                LEFT JOIN Store AS s ON f.store_at = s.s_id
                GROUP BY store_at
                HAVING COUNT(f_id) > 5
                `;
                $.ajax({
                    url: "http://localhost:8789/queryBuilder", 
                    type: "post",
                    data: {
                        sql_query:sqlInput.textContent
                    },
                    success: function(result){
                        console.log(result);
                        buildHtmlTable(JSON.parse(result), "#data-table")
                    }
                });
                break;
            default:
                break;
        }
    }
})
