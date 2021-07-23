<template>
  <div class="container-fluid">
    <div class="card shadow mb-4">
      <div
        class="card-header py-3 d-flex justify-content-between  align-items-center"
      >
        <h6 class="m-0 font-weight-bold text-primary">Quản lý sản phẩm</h6>
        <label
          for="upload"
          class="
                  d-none d-sm-inline-block
                  btn btn-sm btn-primary btn-green
                  shadow-sm
                "
        >
          <i class="fas fa-upload fa-sm text-white-50"></i>Tải lên sản phẩm
          (.csv)</label
        >
        <input
          type="file"
          name="upload"
          id="upload"
          @change="handleFileUpload($event)"
        />
      </div>
      <div class="card-body">
        <div class="table-responsive">
          <table class="table table-bordered" id="dataTable" width="100%">
            <thead>
              <tr>
                <th class="product-id text-center">ID</th>
                <th class="product-name "><span>Tên</span></th>
                <th class="product-price text-center">Giá</th>
                <th class="product-quantity text-center">Số lượng</th>
                <th class="product-sale text-center">Khuyến mãi</th>
                <th class="product-action text-center">Hành động</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="product in products" :key="product.id">
                <td class="product-id text-center">
                  {{ product.id }}
                </td>
                <td class="product-name">
                  <span> {{ product.name }}</span>
                </td>
                <td class="product-price text-center">{{ product.price }}</td>
                <td class="product-quantity text-center">
                  {{ product.quantity }}
                </td>
                <td class="product-sale text-center">
                  {{ product.sale_percent }}%
                </td>
                <td class="product-action text-center">
                  <button class="btn btn-edit">
                    <i class="fas fa-pen"></i>
                  </button>
                  <button class="btn btn-view">
                    <i class="fas fa-eye"></i>
                  </button>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
        <div class="group-btn">
          <button
            class="btn btn-left"
            :disabled="page === 0"
            @click="handleBackPage"
          >
            <i class="fas fa-chevron-left"></i>
          </button>
          <button
            class="btn btn-right"
            @click="handleNextPage"
            :disabled="disabled"
          >
            <i class="fas fa-chevron-right"></i>
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { getProducts } from "../../services/products.service.js";
import { uploadFile } from "../../services/upload.service.js";
export default {
  name: "PageProduct",
  data() {
    return {
      products: [],
      offset: 0,
      limit: 6,
      page: 0,
      disabled: false
    };
  },
  methods: {
    async handleNextPage() {
      this.page++;
      this.offset = this.offset + this.limit;
      let res = await getProducts(this.offset, this.limit);
      this.products = res.data;
      if (!this.products) {
        this.disabled = true;
      }
    },
    async handleBackPage() {
      this.page--;
      this.offset = this.offset - this.limit;
      let res = await getProducts(this.offset, this.limit);
      this.products = res.data;
      if (this.products) {
        this.disabled = false;
      }
    },
    async handleFileUpload(event) {
      const file = event.target.files[0];
      if (file.type === "application/vnd.ms-excel") {
        let formData = new FormData();
        formData.append("file", file);
        const res = await uploadFile(formData);
        let res2 = await getProducts(this.offset, this.limit);
        this.products = res2.data;
        console.log(res);
        alert("Tải lên sản phẩm thành công");
      } else {
        alert("Phải chọn file có định dạng là .csv");
      }
    }
  },
  async created() {
    let res = await getProducts(this.offset, this.limit);
    this.products = res.data;
  }
};
</script>

<style>
td {
  vertical-align: unset !important;
}
.btn-green {
  background-color: green;
  border-color: green;
  margin: 0;
}
.btn-green:hover {
  background-color: green;
  border-color: green;
  opacity: 0.8;
}
.btn-green:focus {
  background-color: green;
  border-color: green;
}
.btn-green:active {
  background-color: green !important;
  border-color: green !important;
}
#upload {
  display: none;
}
.group-btn {
  text-align: right;
}
.btn-left,
.btn-right {
  margin-left: 10px;
  background-color: #4e73df;
  color: white;
}
.btn-left:hover {
  color: white;
  opacity: 0.8;
}
.btn-right:hover {
  color: white;
  opacity: 0.8;
}
.btn-left:disabled {
  opacity: 0.6;
}
.btn-right:disabled {
  opacity: 0.6;
}
.product-name {
  width: 700px;
}
.product-name > span {
  display: -webkit-box;
  -webkit-line-clamp: 1;
  -webkit-box-orient: vertical;
  overflow: hidden;
}
.btn-edit {
  margin-right: 5px;
}
</style>
