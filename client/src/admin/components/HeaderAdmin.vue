<template>
  <nav
    class="
              navbar navbar-expand navbar-light
              bg-white
              topbar
              mb-4
              static-top
              shadow
            "
  >
    <!-- Sidebar Toggle (Topbar) -->
    <button
      id="sidebarToggleTop"
      class="btn btn-link d-md-none rounded-circle mr-3"
    >
      <i class="fa fa-bars"></i>
    </button>

    <!-- Topbar Search -->
    <form
      class="
                d-none d-sm-inline-block
                form-inline
                mr-auto
                ml-md-3
                my-2 my-md-0
                mw-100
                navbar-search
              "
    >
      <div class="input-group">
        <input
          type="text"
          class="form-control bg-light border-0 small"
          placeholder="Search for..."
          aria-label="Search"
          aria-describedby="basic-addon2"
        />
        <div class="input-group-append">
          <button class="btn btn-primary" type="button">
            <i class="fas fa-search fa-sm"></i>
          </button>
        </div>
      </div>
    </form>

    <!-- Topbar Navbar -->
    <ul class="navbar-nav ml-auto">
      <!-- Nav Item - Search Dropdown (Visible Only XS) -->
      <li class="nav-item dropdown no-arrow d-sm-none">
        <a
          class="nav-link dropdown-toggle"
          href="#"
          id="searchDropdown"
          role="button"
          data-toggle="dropdown"
          aria-haspopup="true"
          aria-expanded="false"
        >
          <i class="fas fa-search fa-fw"></i>
        </a>
        <!-- Dropdown - Messages -->
      </li>

      <!-- Nav Item - Alerts -->
      <li class="nav-item dropdown no-arrow mx-1">
        <a
          class="nav-link dropdown-toggle"
          href="#"
          id="alertsDropdown"
          role="button"
          data-toggle="dropdown"
          aria-haspopup="true"
          aria-expanded="false"
        >
          <i class="fas fa-bell fa-fw"></i>
          <!-- Counter - Alerts -->
          <span class="badge badge-danger badge-counter">3+</span>
        </a>
      </li>

      <!-- Nav Item - Messages -->
      <li class="nav-item dropdown no-arrow mx-1">
        <a
          class="nav-link dropdown-toggle"
          href="#"
          id="messagesDropdown"
          role="button"
          data-toggle="dropdown"
          aria-haspopup="true"
          aria-expanded="false"
        >
          <i class="fas fa-envelope fa-fw"></i>
          <!-- Counter - Messages -->
          <span class="badge badge-danger badge-counter">7</span>
        </a>
      </li>

      <div class="topbar-divider d-none d-sm-block"></div>

      <!-- Nav Item - User Information -->
      <li class="nav-item dropdown no-arrow">
        <a
          class="nav-link dropdown-toggle wrap-user"
          href="#"
          id="userDropdown"
          role="button"
          data-toggle="dropdown"
          aria-haspopup="true"
          aria-expanded="false"
        >
          <span class="mr-2 d-none d-lg-inline text-gray-600 small">
            <router-link
              v-if="!$store.getters.isAuthenticated"
              :to="{ name: 'login' }"
            >
              Đăng nhập
            </router-link>
            <span class="username" v-if="$store.getters.isAuthenticated">
              Xin chào, <b>{{ $store.state.user.username }}</b>
              <span @click="handleLogout" class="logout">Đăng xuất</span>
            </span>
          </span>
        </a>
      </li>
    </ul>
  </nav>
</template>

<script>
import { logout } from "../../services/users.service.js";
export default {
  name: "HeaderAdmin",
  methods: {
    handleLogout() {
      logout().then(() => {
        this.$store.dispatch("setUser", null);
        this.$store.commit("setCart", []);
        this.$store.commit("updateLayout", "shop");
        this.$router.push("/login");
      });
    }
  }
};
</script>

<style>
.username {
  position: relative;
}
.logout {
  position: absolute;
  left: 50%;
  transform: translateX(-50%);
  top: 43px;
  width: 150px;
  height: 40px;
  background-color: rgb(248, 248, 248);
  display: flex;
  align-items: center;
  justify-content: center;
  box-shadow: 0px 2px 4px #bfbfbf;
  transition: all 0.2s ease-in-out;
  opacity: 0;
  visibility: hidden;
}
.wrap-user {
  cursor: pointer;
}
.wrap-user:hover .logout {
  opacity: 1;
  visibility: visible;
}
</style>
