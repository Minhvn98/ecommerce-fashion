# Bài tập cuối khóa Talent Fresher FullStack Developer tại OCG.
Đề tài: Shop bán quần áo online

## 1. Repository's info

| Fullname     | Class    |
| ------------ | -------- |
| Lê Văn Khánh | TFS - 01 |
| Vũ Ngọc Minh | TFS - 01 |

## 2. Directory's tree

```
.
├── client
│   ├── public
│   └── src
│       ├── admin
│       │   ├── components
│       │   └── pages
│       ├── assets
│       │   └── styles
│       ├── components
│       ├── layout
│       ├── pages
│       ├── router
│       ├── services
│       ├── store
│       │   └── modules
│       └── utils
├── elasticsearch
│   ├── controllers
│   ├── models
│   ├── pkg
│   ├── repository
│   ├── routes
│   └── tmp
└── server
    ├── config
    ├── controllers
    ├── database
    ├── middlewares
    ├── models
    ├── public
    │   └── uploads
    │       └── images
    ├── repository
    ├── router
    ├── sql
    └── utils

```
<br/>

<br/>

## 3. Hướng dẫn run project: 
>git clone https://github.com/Minhvn98/ecommerce-fashion.git
<br />

### Run frontend:
>1. cd client
>2. yarn install
>3. yarn serve
<br />

### Run backend:
>1. cd server
>2. go mod tidy
>3. go run main.go (nếu cài air thì chạy air cho nhanh:)))
(Nếu muốn chạy chức năng thanh toán thì cài thêm ngrok nữa nhé nhưng cái này cũng hay lỗi lắm :))
(Tạo database trước nhé trong folder sql có sẵn các câu lệnh rồi)
<br />

### Run elasticsearch
>1. Cài docker compose ở dưới phần resource (nhớ run  trước).
>2. cd elasticsearch
>3. go run main.go (hoặc dùng air)

## 4. Usecase diagram: <br />

![image](https://user-images.githubusercontent.com/37922407/126858787-b9dea94a-373a-41d8-8a02-bef603f447a0.png)<br/>

## 5. Database diagram: <br />

![image](https://user-images.githubusercontent.com/37922407/126858719-1a73e4d8-749e-4c49-ac87-b30643544e54.png)


## 6. Một số hình ảnh trong project:
![image](https://user-images.githubusercontent.com/37922407/126858597-723e04ea-7850-4a13-ab19-690338f9bc93.png)
<br/>
![image](https://user-images.githubusercontent.com/37922407/126858602-12853fec-9377-414b-9f9a-9392725674a8.png)
<br/>
![image](https://user-images.githubusercontent.com/37922407/126858609-84a4a03b-de61-48a9-b3de-551797264874.png)
<br/>
![image](https://user-images.githubusercontent.com/37922407/126858621-eb112097-99e0-4f98-bc61-9d0cf1e41cb6.png)
<br/>
![image](https://user-images.githubusercontent.com/37922407/126858629-ffbff18c-5cd2-437d-9af8-3be2ed9ce3e8.png)
<br/>
![image](https://user-images.githubusercontent.com/37922407/126858645-c72a5db3-7878-48c9-9d31-f739ce205b79.png)
<br/>
![image](https://user-images.githubusercontent.com/37922407/126858657-37de9643-976e-4c23-811b-5bb454412ad6.png)
<br/>
...

## 7. Resource:
- Slide: https://docs.google.com/presentation/d/12EEs0fhV96Zvl4gqpXbSKsUZxAfmsWXrefvq-sqhoeM

- Usecase: https://drive.google.com/file/d/1PmDolVZHcW0IPBj3S8h9FAABKNm1kTq-/view?usp=sharing

... <br />
Bye!

