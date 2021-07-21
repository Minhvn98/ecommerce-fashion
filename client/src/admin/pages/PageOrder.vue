<template>
  <div class="container-fluid">
    <div class="card shadow mb-4">
      <div class="card-header py-3">
        <h6 class="m-0 font-weight-bold text-primary">Quản lý đơn hàng</h6>
      </div>
      <div class="card-body">
        <div class="table-responsive">
          <table class="table table-bordered" id="dataTable" width="100%">
            <thead>
              <tr>
                <th class="product-id">ID</th>
                <th class="product-name">Tên</th>
                <th class="product-price">Giá</th>
                <th class="product-quantity">Số lượng</th>
                <th class="product-sale">Khuyến mãi</th>
                <th class="product-action">Hành động</th>
              </tr>
            </thead>
            <tbody>
              <tr>
                <td class="product-id">1</td>
                <td class="product-name">
                  Support Customer
                </td>
                <td class="product-price">New York</td>
                <td class="product-quantity">27</td>
                <td class="product-sale">2011/01/25</td>
                <td class="product-action">$112,000</td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "PageOrder"
};
</script>

<style></style>
