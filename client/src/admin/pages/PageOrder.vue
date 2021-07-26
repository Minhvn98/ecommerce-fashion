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
                <th class="order-id">ID</th>
                <th class="order-name">Tên</th>
                <th class="order-phone">SDT</th>
                <th class="order-location">Địa chỉ</th>
                <th class="order-payment">Thanh toán</th>
                <th class="order-status">Trạng thái</th>
                <th class="order-created">Ngày tạo</th>
                <th class="order-action">Hành động</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="order in orders" :key="order.id">
                <td class="order-id">{{ order.id }}</td>
                <td class="order-name">
                  {{ order.first_name }} {{ last_name }}
                </td>
                <td class="order-phone">{{ order.phone }}</td>
                <th class="order-location">{{ order.location }}</th>
                <td class="order-payment">{{ order.payment }}</td>
                <td class="order-status">{{ order.status }}</td>
                <td class="order-created">{{ order.created_at }}</td>
                <td class="order-action text-center">
                  <button
                    @click="handleOpenModal(order.id, order.status)"
                    class="btn btn-edit"
                  >
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
      </div>
    </div>
  </div>
  <ModalSmaill
    :isShow="isOpenModal"
    :id="idOrder"
    @hide="hideModel"
    :statusOrder="statusOrder"
    @close="closeModal"
    @submit="updateStatus"
  />
</template>

<script>
import {
  getAllOrder,
  updateStatusOrder,
} from "../../services/order.service.js";

import ModalSmaill from "../components/ModalSmall.vue";

export default {
  name: "PageOrder",
  components: {
    ModalSmaill,
  },
  data() {
    return {
      orders: null,
      idOrder: 0,
      statusOrder: "",
      isOpenModal: false,
    };
  },
  methods: {
    handleOpenModal(id, statusOrder) {
      // if (statusOrder === "Chờ thanh toán") {
      this.idOrder = id;
      this.statusOrder = statusOrder;
      this.isOpenModal = true;
      // } else {
      // alert("Đơn hàng đã được hoàn thành");
      // }
    },
    hideModel() {
      this.isOpenModal = false;
    },
    updateStatus(status) {
      updateStatusOrder(this.idOrder, status).then(async () => {
        this.isOpenModal = false;
        alert("Cập nhập trạng thái thành công");
        let res = await getAllOrder();
        this.orders = res.data;
      });
    },
    closeModal() {
      this.isOpenModal = false;
    },
  },
  async created() {
    let res = await getAllOrder();
    this.orders = res.data;
  },
};
</script>

