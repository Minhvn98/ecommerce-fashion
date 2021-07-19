<template>
  <div class="container">
    <div class="bill-wrapper">
      <h1 class="title">Hóa Đơn</h1>
      <div class="content">
        <div class="info">
          <div class="info-customer">
            <ul>
              <li class="item"><strong>Khách Hàng:</strong> Lê Văn Khánh</li>
              <li class="item">
                <strong>Địa Chỉ:</strong> Đường Láng, Đống Đa, Hà Nội
              </li>
              <li class="item"><strong>SĐT:</strong> 038.123.123</li>
            </ul>
          </div>

          <div class="info-bill">
            <ul>
              <li class="item"><strong>Hóa Đơn:</strong> 180721</li>
              <li class="item"><strong>Ngày Thanh Toán:</strong> 18/07/2021</li>
              <li class="item"><strong>Trạng Thái:</strong> Đã thanh toán</li>
              <h1 class="total">Tổng tiền: 690.000 VND</h1>
            </ul>
          </div>
        </div>

        <div class="bill-detail">
          <!--                <h3 class="title">Chi tiết</h3>-->
          <div class="product">
            <div class="product-image">
              <img src="../assets/anh2.jpeg" alt="" />
            </div>
            <div class="product-name">
              Đồ bộ Yukata hình voi JANE màu xám Y04
            </div>
            <div class="product-price">95.000 VND</div>
          </div>

          <div class="product">
            <div class="product-image">
              <img src="../assets/anh2.jpeg" alt="" />
            </div>
            <div class="product-name">
              Đồ bộ Yukata hình voi JANE màu xám Y04
            </div>
            <div class="product-price">95.000 VND</div>
          </div>

          <div class="product">
            <div class="product-image">
              <img src="../assets/anh2.jpeg" alt="" />
            </div>
            <div class="product-name">
              Đồ bộ Yukata hình voi JANE màu xám Y04
            </div>
            <div class="product-price">95.000 VND</div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "BillPage",
};
</script>

<style scoped>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

.total {
  font-size: 22px;
  font-weight: 400;
  margin-top: 20px;
  border-top: 1px solid #ddd;
  padding: 12px 0;
  color: #ee4d2d;
}

strong {
  color: #ee4d2d;
}

.bill-detail {
  width: 70%;
  padding: 15px;
}

.info {
  width: 30%;
  padding: 15px;
  background: #f7f7f7;
}

.content {
  display: flex;
  margin-top: 30px;
  background: #fff;
  /* padding: 15px; */
  border-radius: 3px;
}

.item {
  list-style: none;
  padding: 5px 0;
}

body {
  background: #e9e9e9;
}

.product-price {
  width: 17%;
  color: #ee4d2d;
  display: flex;
  align-items: center;
}

.product-name {
  width: 70%;
}

.product-image img {
  width: 100%;
  border-radius: 3px;
}

.product-image {
  width: 12%;
  margin-right: 15px;
}

.product {
  display: flex;
  background: #fff;
  padding: 10px;
  border-radius: 3px;
  margin: 10px 0;
  border: 1px solid #ddd;
}

.title {
  font-size: 23px;
  font-weight: 400;
  color: #ee4d2d;
}

.container {
  width: 80%;
  margin: auto;
}

.bill-wrapper {
  margin-top: 115px;
  width: 100%;
}
</style>
