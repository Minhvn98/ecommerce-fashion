<template>
  <div class="modal">
    <div class="modal-dialog">
      <div class="modal-content">
        <div class="modal-header">
          <h5 class="modal-title">Bạn có chắc chắc muốn xóa sản phẩm này ?</h5>
          <button type="button" class="btn-close" @click="confirmNo">X</button>
        </div>
        <div class="modal-body">
          <p>
            {{ nameProduct }}
          </p>
        </div>
        <div class="modal-footer">
          <button type="button" class="btn btn-ok" @click="confirmYes">
            Có
          </button>

          <button type="button" class="btn btn-no" @click="confirmNo">
            Không
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "ModalDelete",
  props: {
    nameProduct: String,
  },

  methods: {
    confirmYes() {
      this.$emit("confirm", true);
    },

    confirmNo() {
      this.$emit("confirm", false);
    },
  },
};
</script>

<style scoped>
.modal {
  display: flex;
  justify-content: center;
  align-items: center;
}

.modal-dialog {
  max-width: 540px;
  position: fixed;
  top: 100px;
  background: rgb(255, 255, 255);
  margin: auto;
  padding: 15px;
  border-radius: 5px;
  z-index: 999;
}

.modal-header {
  display: flex;
  justify-content: space-between;
  border-bottom: 1px solid #ececec;
  padding-bottom: 15px;
  align-items: center;
}
.modal-title {
  font-weight: 500;
  font-size: 22px;
  color: #ee4d2d;
  margin: 10px 0;
}
.btn-close {
  width: 40px;
  height: 30px;
  border: none;
  cursor: pointer;
  font-size: 20px;
  background: transparent;
  font-weight: 600;
  color: #555;
}
.btn-close:hover {
  color: #000;
}

.modal-footer {
  padding-top: 25px;
  border-top: 1px solid rgb(238, 237, 237);
}

.modal-body {
  padding-top: 25px;
}
.modal-body p {
  padding-bottom: 25px;
  font-size: 17px;
}

.btn {
  padding: 11px 20px;
  font-size: 18px;
  width: calc(50% - 5px);
  cursor: pointer;
  border: none;
  border-radius: 4px;
}
.btn.btn-ok {
  background: #ee4d2d;
  color: #fff;
  margin-right: 10px;
}

.btn-no {
  background-color: #fff;
  border: 1px solid #ddd;
}

.btn-ok:hover {
  background-color: #d6563d;
}

.btn-no:hover {
  background-color: #eee;
}
.background-black {
  position: fixed;
  top: 0;
  z-index: 9;
  width: 100%;
  height: 900px;
  background: rgba(0, 0, 0, 0.507);
}
</style>