<template>
  <div v-show="isShow" class="wrap-modal">
    <div class="modal-status">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title">Cập nhập trạng thái đơn hàng: {{ id }}</h5>
          </div>
          <div class="modal-body">
            <p>Trạng thái</p>
            <select @change="handleSelect($event)">
              <option value="0">Đã thanh toán</option>
              <option value="1">Thanh toán thất bại</option>
            </select>
          </div>
          <div class="modal-footer">
            <button type="button" class="btn btn-primary" @click="updateStatus">
              Cập nhập
            </button>
            <button type="button" class="btn " @click="hideModal">
              Hủy bỏ
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
  <div v-show="isShow" class="overlay"></div>
</template>

<script>
export default {
  name: "ModalSmaill",
  props: {
    isShow: Boolean,
    id: Number,
    statusOrder: String
  },
  data() {
    return {
      status: "Đã thanh toán"
    };
  },
  methods: {
    HandleHideModal() {
      this.$emit("hide");
    },
    handleSelect(e) {
      let value = e.target.value;
      if (value == 0) {
        this.status = "Đã thanh toán";
      }
      if (value == 1) {
        this.status = "Thanh toán thất bại";
      }
    },
    updateStatus() {
      this.$emit("submit", this.status);
    },
    hideModal() {
      this.$emit("close");
    }
  }
};
</script>

<style>
select {
  width: 100%;
  height: 50px;
}
.wrap-modal {
  position: fixed;
  top: 0;
  left: 0;
  bottom: 0;
  left: 0;
  width: 100vw;
  height: 100vh;
  background-color: transparent;
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
}
.overlay {
  position: fixed;
  top: 0;
  left: 0;
  bottom: 0;
  left: 0;
  width: 100vw;
  height: 100vh;
  background-color: rgba(0, 0, 0, 0.5);
  z-index: 999;
}
.modal-status {
  width: 400px;
  height: 300px;
}
</style>
