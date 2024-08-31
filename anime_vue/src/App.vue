<template>
    <div class="mainBody">
        <Header class="header"></Header>
        <br>
        <router-view></router-view>
    </div>
</template>

<script setup>
import Header from './components/HeaderPage.vue';
</script>

<style scoped>
.mainBody {
    display: flex;
    justify-content: center;
    flex-direction: column;
    align-items: center;
    width: 100%;
    height: 100%;
}

.header {
    padding: 10px;
}
</style>
