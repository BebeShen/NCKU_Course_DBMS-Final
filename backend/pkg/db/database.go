package db

import (
	"fmt"
	"database/sql"

    // . "github.com/bebeshen/efrs/pkg/utils"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Connect() () {
    // global db connection object
	DB, _ = sql.Open("sqlite3", "../db/sqlite/efrs.db")
}

func CreateTable() {
    /* Customer Table */
    fmt.Println("create Customer table")
    customer_table := `
    CREATE TABLE IF NOT EXISTS Customer(
        c_id INTEGER PRIMARY KEY AUTOINCREMENT,
        username TEXT NULL,
        password TEXT NULL,
        c_location TEXT NULL
    );
	`
	_, err := DB.Exec(customer_table)
    if (err!=nil) {
        fmt.Println("Fail creating Customer table",err)
    }

    /* Employee Table */
    fmt.Println("create Employee table")
    employee_table := `
    CREATE TABLE IF NOT EXISTS Employee(
        e_id INTEGER PRIMARY KEY AUTOINCREMENT,
        username TEXT NULL,
        password TEXT NULL,
        work_for INTEGER
    );
	`
	_, err = DB.Exec(employee_table)
    if (err != nil) {
        fmt.Println("Fail creating Employee table", err)
    }

    /* Store Table */
    fmt.Println("create Store table")
    store_table := `
    CREATE TABLE IF NOT EXISTS Store(
        s_id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NULL,
        location TEXT NULL,
        type TEXT NULL,
        owner INTEGER
    );
	`
	_, err = DB.Exec(store_table)
    if (err != nil) {
        fmt.Println("Fail creating Store table", err)
    }

    /* Food Table */
    fmt.Println("create Food table")
    food_table := `
    CREATE TABLE IF NOT EXISTS Food(
        f_id INTEGER PRIMARY KEY AUTOINCREMENT,
        category TEXT NULL,
        name TEXT NULL,
        expireDate TEXT NULL,
        price REAL NULL,
        discount REAL DEFAULT 0.0,
        store_at INTEGER
    );
	`
	_, err = DB.Exec(food_table)
    if (err != nil) {
        fmt.Println("Fail creating Food table", err)
    }

    /* Wasted Table */
    fmt.Println("create Wasted table")
    wasted_table := `
    CREATE TABLE IF NOT EXISTS Wasted(
        w_id INTEGER PRIMARY KEY AUTOINCREMENT,
        f_id INTEGER,
        category TEXT NULL,
        name TEXT NULL
    );
	`
	_, err = DB.Exec(wasted_table)
    if (err != nil) {
        fmt.Println("Fail creating Wasted table", err)
    }

    /* Order Table */
    fmt.Println("create Orders table")
    // use orders to prevent conflict against sql keyword
    orders_table := `
    CREATE TABLE IF NOT EXISTS Orders(
        c_id INTEGER NOT NULL,
        f_id INTEGER NOT NULL,
        s_id INTEGER NOT NULL,
        message TEXT NULL,
        status TEXT NULL,
        PRIMARY KEY (c_id, f_id)
    );
	`
	_, err = DB.Exec(orders_table)
    if (err != nil) {
        fmt.Println("Fail creating Orders table", err)
    }
}

func QueryBuilder(sql_query string) (*sql.Rows, error) {
    fmt.Println("SQL Query:", sql_query)
    rows, err := DB.Query(sql_query)
    if (err!=nil) {
        fmt.Println(err)
        return nil, err
    }

    cols, _ := rows.Columns()
    fmt.Println("Columns: ")
	for i := range cols {
		fmt.Print(cols[i])
		fmt.Print("\t")
	}
    fmt.Println()
    
    return rows, nil
}

func InsertDefaultData()  {
    // customer
    customer_data := `
        INSERT INTO Customer (username, password, c_location) 
        VALUES 
            ('餓死鬼1號', 'psw1', '701台南市東區莊敬里 中華東路一段 66號'),
            ('小赤佬愛好者', 'psw2', '701台南市東區青年路416-3號'),
            ('小紅帽老闆', 'psw3', '701台南市東區勝利路159號'),
            ('4號打擊手', 'psw3', '701台南市東區育樂街160號'),
            ('5索', 'psw3', '701台南市東區大學路26號'),
            ('6中萌', 'psw3', '701台南市東區大學路3號'),
            ('77志棋', 'psw3', '701台南市東區大學路1號'),
            ('8+9', 'psw3', '710台南市永康區中華路34號'),
            ('9鬼波吉', 'psw3', '704台南市北區東豐路257號'),
            ('十十十十', 'psw3', '701台南市東區中華東路一段366號')
            ;
    `
	if _, err := DB.Exec(customer_data) ; err != nil {
        fmt.Println("insert default customer fail", err)
    }
    // store
    store_data := `
        INSERT INTO Store (location, name, type, owner) 
        VALUES 
            ('701台南市東區中華東路一段24號', '7-ELEVEN 一心門市','7-11', 1),
            ('701台南市東區大學路1號', '全家便利商店-台南成大店', '全家', 4),
            ('704台南市北區開元路212巷150號', '全家便利商店 台南三順店', '全家', 6),
            ('701台南市東區東安路291號293號1樓', '7-ELEVEN 東和門市', '7-11', 8),
            ('701台南市東區光明街217號219號', '全家便利商店-台南東平店', '全家', 10),
            ('704台南市北區南園街102號號1+2樓', '全家便利商店-台南南園店', '全家', 12),
            ('700台南市中西區中山路207號', '全家便利商店 台南前站店', '全家', 14),
            ('704台南市北區小東路307巷1樓 75號77號', '7-ELEVEN 鈺勝門市', '7-11', 16),
            ('701台南市東區裕農路888號', '7-ELEVEN 裕孝門市', '7-11', 18),
            ('710台南市永康區復興路405號', '萊爾富便利商店 永康榮醫店', '萊爾富', 20)
            ;
    `
	if _, err := DB.Exec(store_data) ; err != nil {
        fmt.Println("insert default store fail", err)
    }
    // employee
    employee_data := `
        INSERT INTO Employee (username, password, work_for) 
        VALUES 
            ('老闆1', 'psw1', 1),
            ('員工1_1', 'psw2', 1),
            ('員工1_2', 'psw3', 1),
            ('老闆2', 'psw4', 2),
            ('員工2_1', 'psw5', 2),
            ('老闆3', 'psw6', 3),
            ('員工3_1', 'psw7', 3),
            ('老闆4', 'psw8', 4),
            ('員工4_1', 'psw9', 4),
            ('老闆5', 'psw4', 5),
            ('員工5_1', 'psw5', 5),
            ('老闆6', 'psw6', 6),
            ('員工6_1', 'psw7', 6),
            ('老闆7', 'psw8', 7),
            ('員工7_1', 'psw9', 7),
            ('老闆8', 'psw4', 8),
            ('員工8_1', 'psw5', 8),
            ('老闆9', 'psw6', 9),
            ('員工9_1', 'psw7', 9),
            ('老闆10', 'psw8', 10),
            ('員工10_1', 'psw9', 10)
            ;
    `
	if _, err := DB.Exec(employee_data) ; err != nil {
        fmt.Println("insert default employee fail", err)
    }
    // food
    food_data := `
    INSERT INTO Food(category, name, expireDate, price, discount, store_at) 
    VALUES 
        ('riceroll', '炙燒明太子鮭魚飯糰', '2022-06-12', 33, 0, 1),
        ('riceroll', '雞肉飯飯糰', '2022-06-12', 28, 0, 1),
        ('luwei', '石安牧場溫泉蛋', '2022-06-12', 46, 0, 1),
        ('veg', '松露風味野菇烘蛋捲餅', '2022-06-12', 52, 0, 1),
        ('bread', '煉乳牛奶麵包', '2022-06-12', 35, 0, 1),
        ('bread', '檸檬雪片天使蛋糕', '2022-06-12', 45, 0, 1),
        ('riceroll', '哇沙米鮭魚飯糰', '2022-06-13', 33, 0, 1),
        ('riceroll', '握便當-經典奮起湖', '2022-06-13', 45, 0, 1),
        ('ohlala', '雙拼起司奶香焗飯', '2022-06-13', 65, 0, 1),
        ('ohlala', '雙拼起司奶香焗飯', '2022-06-14', 65, 0, 1),
        ('luwei', '石安牧場溫泉蛋', '2022-06-12', 46, 0, 1),
        ('ohlala', '培根奶油風味義大利麵', '2022-06-14', 65, 0, 1),
        ('riceroll', '請客樓-麻油雞飯糰', '2022-06-14', 49, 0, 1),
        ('riceroll', '炙燒雪花牛飯糰', '2022-06-15', 36, 0, 1),
        ('riceroll', '魚卵小龍蝦沙拉飯糰', '2022-06-15', 33, 0, 2),
        ('riceroll', '御選肉鬆飯糰', '2022-06-15', 28, 0, 1),
        ('riceroll', '炙燒明太子鮭魚飯糰', '2022-06-15', 33, 0, 1),
        ('veg', '松露風味野菇烘蛋捲餅', '2022-06-16', 52, 0, 1)
        ;
    `
    if _, err := DB.Exec(food_data) ; err != nil {
        fmt.Println("insert default food fail", err)
    }
    // wasted
    wasted_data := `
    INSERT INTO Wasted(f_id, category, name) 
    VALUES 
        (1, 'riceroll', '炙燒明太子鮭魚飯糰'),
        (2, 'riceroll', '雞肉飯飯糰'),
        (3, 'luwei', '石安牧場溫泉蛋'),
        (4, 'veg', '松露風味野菇烘蛋捲餅'),
        (5, 'bread', '煉乳牛奶麵包'),
        (6, 'bread', '檸檬雪片天使蛋糕'),
        (7, 'riceroll', '哇沙米鮭魚飯糰'),
        (8, 'riceroll', '握便當-經典奮起湖'),
        (9, 'ohlala', '雙拼起司奶香焗飯'),
        (10, 'ohlala', '雙拼起司奶香焗飯')
        ;
    `
    if _, err := DB.Exec(wasted_data) ; err != nil {
        fmt.Println("insert default Wasted fail", err)
    }
    // order
    order_data := `
    INSERT INTO Orders(c_id, f_id, s_id, message, status) 
    VALUES 
        (1, 1, 1, '我是省錢人', '預約'),
        (1, 2, 1, '', '預約'),
        (1, 3, 1, '', '預約'),
        (1, 4, 1, '', '預約'),
        (1, 5, 1, '', '預約'),
        (1, 6, 1, '', '預約'),
        (1, 7, 1, '', '預約'),
        (1, 8, 1, '', '預約'),
        (1, 9, 1, '', '預約'),
        (1, 10, 1, '', '預約')
        ;
    `
    if _, err := DB.Exec(order_data) ; err != nil {
        fmt.Println("insert default order fail", err)
    }
}