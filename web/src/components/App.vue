<template>
  <div class="flexible-content">
    <!--Navbar-->
    <mdb-navbar class="flexible-navbar white" light position="top" scrolling>
      <mdb-navbar-brand href="/">
        Cerberus
      </mdb-navbar-brand>
      <mdb-navbar-toggler>
        <mdb-navbar-nav left>
          <mdb-nav-item to="/" waves-fixed active class="active">
            Home
          </mdb-nav-item>
        </mdb-navbar-nav>
      </mdb-navbar-toggler>
    </mdb-navbar>
    <!--/.Navbar-->
    <!-- Sidebar -->
    <div class="sidebar-fixed position-fixed">
      <router-link to="/"
        ><a class="logo-wrapper"
          ><img alt="" class="img-fluid" src="@/assets/logo_360.png"/></a
      ></router-link>
      <mdb-list-group class="list-group-flush">
        <router-link to="/" @click.native="activeItem = 1">
          <mdb-list-group-item
            :action="true"
            :class="activeItem === 1 && 'active'"
          >
            <mdb-icon icon="chart-pie" class="mr-3" />Dashboard
          </mdb-list-group-item>
        </router-link>
        <router-link to="/services" @click.native="activeItem = 3">
          <mdb-list-group-item
            :action="true"
            :class="activeItem === 3 && 'active'"
          >
            <mdb-icon icon="cogs" class="mr-3" />Services
          </mdb-list-group-item>
        </router-link>
        <router-link to="/maps" @click.native="activeItem = 4">
          <mdb-list-group-item
            :action="true"
            :class="activeItem === 4 && 'active'"
          >
            <mdb-icon icon="users" class="mr-3" />Users
          </mdb-list-group-item>
        </router-link>
        <router-link to="/profile" @click.native="activeItem = 2">
          <mdb-list-group-item
            :action="true"
            :class="activeItem === 2 && 'active'"
          >
            <mdb-icon icon="user" class="mr-3" /> My Account
          </mdb-list-group-item>
        </router-link>
      </mdb-list-group>
    </div>
    <!-- /Sidebar  -->
    <main>
      <div class="mt-5 p-5">
        <router-view />
      </div>
      <ftr color="primary-color-dark" class="text-center font-small darken-2">
        <p class="mb-0 py-3 text-center">
          &copy; {{ new Date().getFullYear() }} Copyright:
          <a href="https://www.forsam.education"> ForSAM Education </a>
          <a href="https://github.com/forsam-education/cerberus"
            ><mdb-icon fab icon="github" class="mr-3"
          /></a>
        </p>
        <p class="footer-copyright mb-0 py-3 text-center">
          Template by <a href="https://mdbootstrap.com/"> MDBootstrap </a>
        </p>
      </ftr>
    </main>
  </div>
</template>

<script>
import {
  mdbNavbar,
  mdbNavbarBrand,
  mdbNavItem,
  mdbNavbarNav,
  mdbNavbarToggler,
  mdbIcon,
  mdbListGroup,
  mdbListGroupItem,
  mdbFooter,
  waves,
} from 'mdbvue';

export default {
  name: 'AdminTemplate',
  components: {
    mdbNavbar,
    mdbNavbarBrand,
    mdbNavItem,
    mdbNavbarNav,
    mdbNavbarToggler,
    mdbListGroup,
    mdbListGroupItem,
    mdbIcon,
    ftr: mdbFooter,
  },
  data () {
    return {
      activeItem: 1,
    };
  },
  beforeMount () {
    this.activeItem = this.$route.matched[0].props.default.page;
  },
  mixins: [waves],
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
main {
  background-color: #ededee;
}

.flexible-content {
  transition: padding-left 0.3s;
  padding-left: 270px;
}

.flexible-navbar {
  transition: padding-left 0.5s;
  padding-left: 270px;
}

.sidebar-fixed {
  left: 0;
  top: 0;
  height: 100vh;
  width: 270px;
  box-shadow: 0 2px 5px 0 rgba(0, 0, 0, 0.16), 0 2px 10px 0 rgba(0, 0, 0, 0.12);
  z-index: 1050;
  background-color: #fff;
  padding: 1.5rem;
  padding-top: 0;
}

.sidebar-fixed .logo-wrapper img {
  width: 100%;
  padding: 2.5rem;
}

.sidebar-fixed .list-group-item {
  display: block !important;
  transition: background-color 0.3s;
}

.sidebar-fixed .list-group .active {
  background-color: #464d5d;
  border-color: #464d5d;
  box-shadow: 0 2px 5px 0 rgba(0, 0, 0, 0.16), 0 2px 10px 0 rgba(0, 0, 0, 0.12);
  border-radius: 5px;
}

@media (max-width: 1199.98px) {
  .sidebar-fixed {
    display: none;
  }
  .flexible-content {
    padding-left: 0;
  }
  .flexible-navbar {
    padding-left: 10px;
  }
}
</style>
<style>
.navbar-light .navbar-brand {
  margin-left: 15px;
  color: #606879 !important;
  font-weight: bolder;
}
</style>
