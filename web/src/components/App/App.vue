<template>
  <v-app>
    <router-view />
    <Footer />
  </v-app>
</template>

<script>
import Footer from '@/components/Footer/Footer.vue';
import '@MasterStyle';

export default {
  name: 'App',
  components: { Footer },
};
</script>
