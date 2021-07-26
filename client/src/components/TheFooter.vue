<template>
  <footer class="footer-wrap">
    <div class="footer">
      <div class="footer-item">
        <h4 class="footer-title">Về Shopee</h4>
        <ul>
          <li>Giới Thiệu Về Shopee Việt Nam</li>
          <li>Tuyển Dụng</li>
          <li>Điều Khoản Shoppe</li>
          <li>Chính Hãng</li>
          <li>Kênh Người Bán</li>
          <li>Flash Sale</li>
          <li>Liên hệ với truyền thông</li>
        </ul>
      </div>

      <div class="footer-item">
        <h4 class="footer-title">Thanh toán</h4>
        <ul>
          <li><i class="fab fa-cc-paypal"></i> Paypal</li>
          <li><i class="fab fa-cc-visa"></i> Visa</li>
          <li><i class="fab fa-cc-jcb"></i> JCB</li>
          <li><i class="far fa-credit-card"></i> Credit Card</li>
          <li><i class="far fa-money-bill-alt"></i> Ví Shopee</li>
        </ul>
      </div>
      <div class="footer-item">
        <h4 class="footer-title">Theo dõi chúng tôi trên</h4>
        <ul>
          <li><i class="fab fa-facebook-f"></i> Facebook</li>
          <li><i class="fab fa-linkedin"></i> LinkedIn</li>
          <li><i class="fab fa-instagram"></i> Instagram</li>
        </ul>
      </div>
      <div class="footer-item">
        <h4 class="footer-title">Tải ứng dụng ngay thôi</h4>
        <div class="footer-app">
          <img class="qrcode" src="../assets/qrcode.png" alt="" />
        </div>
      </div>
    </div>
  </footer>
</template>

<script>
export default {
  name: "TheFooter",
};
</script>

<style scoped>
/* footer  */
.footer-wrap {
  width: 100%;
  background: #f5f5f5;
  color: #8c8787;
  padding-bottom: 30px;
  margin-top: 10px;
}

.footer {
  display: flex;
  width: calc(80% - 30px);
  margin: auto;
}

.footer-item {
  width: 25%;
}

.footer-item li {
  list-style: none;
  padding: 10px 0;
  cursor: pointer;
}

.footer-item li:hover {
  color: var(--primary-color);
}

.footer-title {
  padding: 20px 0;
  font-size: 16px;
  color: #333;
}

.footer-item i {
  color: #333;
  margin-right: 15px;
}

.footer-app {
  display: flex;
}

.qrcode {
  max-width: 60%;
}

.banner-wrap {
  margin-top: 115px !important;
}

@media screen and (max-width: 576px) {
  .footer {
    width: calc(100% - 30px);
    flex-wrap: wrap;
  }

  .footer-item {
    width: 50%;
  }
}
</style>
