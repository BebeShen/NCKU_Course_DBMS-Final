# DBMS Final Project

資料庫系統導論期末 Project: Expired Food Reservation System (EFRS)

## 系統簡介



## 系統架構與環境

- Programming Language：Golang
- Database：SQLite3
- Environment：Windows 10

## 介面截圖

![screenshot 1](./images/Customer_select_embed.jpg)

![screenshot 2](./images/Customer_select_query.jpg)

![screenshot 3](./images/Food_insert_embed.jpg)

## 使用說明

Prerequisites

- Golang

1. Run backend server

    ```shell
    cd backend

    go run .
    ```
2. Open frontend(`index.html`)

## 資料庫設計

### ER Diagram

![ER Diagram](./images/EFRS_ERmodel.drawio.png)

### Relation Schema

![Relation Schema](./images/RelationSchema.png)

### Table、Attribute、Relationship 說明

- **`Customer`**：紀錄本系統註冊之顧客資訊。
  - `c_id`：**Primary Key**，customer的id(unique)。
  - `username`：顧客註冊帳號。
  - `password`：顧客註冊密碼。
  - `c_location`：顧客當前位置。
- **`Employee`**：紀錄每家商店的員工。
  - `e_id`：**Primary Key**，employee的id(unique)。
  - `username`：員工帳號。
  - `password`：員工密碼。
  - `work_for`:在哪間商店上班->`Store.s_id`。
- **`Store`**：紀錄商店資訊。
  - `s_id`：**Primary Key**，商店的id(unique)。
  - `s_location`：商店的位址。
  - `name`：商店名稱。
  - `type`：商店種類，例如7-11、全家、麵包店...。
  - `owner`：經理是哪位->`Employee.e_id`。
- **`Food`**：紀錄所有登入之食物。
  - `f_id`：**Primary Key**，食物的id(unique)。
  - `category`：食物的種類，例如飯糰、麵包...。
  - `name`：食物的名稱。
  - `expireDate`：食物有效期限。
  - `price`：販售價格。
  - `discount`：即期優惠。
  - `store_at`：販售在哪間商店->`Store.s_id`。
- **`Wasted`**：紀錄所有被浪費掉的食物。
  - `w_id`：**Primary Key**，Wasted的id(unique)。
  - `f_id`：Food的id(unique)->`Food.f_id`。
  - `category`：食物的種類，例如飯糰、麵包...。
  - `name`：食物的名稱。
- **`Orders`**：即期食品的預約訂單。
  - `c_id`：**Primary Key**，顧客的id(unique)->`Customer.c_id`。
  - `f_id`：**Primary Key**，食物的id(unique)->`Food.f_id`。
  - `s_id`：商店的id(unique)->->`Store.s_id`。
    - `(c_id,f_id)`為composite PK。
  - `message`：顧客於此訂單之留言。
  - `status`：此訂單交易狀態，`預約`、`取消`、`結單`。

### Embedded SQL 說明

1. **SELECT-FROM-WHERE**
   - Customer：查詢所有顧客資訊
  
    ```plain
    SELECT * FROM Customer
    ```

   - Employee：查詢所有員工資訊
  
    ```plain
    SELECT * FROM Employee
    ```

   - Food：查詢所有食物資訊
  
    ```plain
    SELECT * FROM Food
    ```

2. **DELETE**
    - Customer：顧客刪除帳號

    ```plain
    DELETE FROM Customer WHERE c_id=1;
    ```

3. **INSERT**
   - Customer：新顧客註冊

    ```plain
    INSERT INTO customer (username, password, c_location) VALUES ('test', 'test', '701台南市東區莊敬里 中華東路一段 66號');
    ```

4. **UPDATE**
   - Customer：顧客更新當前GPS位置

    ```plain
    UPDATE Customer SET c_location='701台南市東區大學路1號' WHERE c_id = 1;
    ```
   
5. **EXISTS、NOT EXISTS**
   
6. **AGGREGATE**
