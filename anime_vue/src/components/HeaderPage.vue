<template>
  <header class="header">
    <div class="left-section">
      <img class="avatar" :src="avatarUrl" alt="Avatar" />
      <a :href="githubLink" target="_blank" class="github-button">GitHub</a>
    </div>
    <button class="logout-button" @click="logout">登出</button>
  </header>
</template>

<script setup>
import { useRouter } from 'vue-router';

const avatarUrl = 'https://avatars.githubusercontent.com/u/93471086?s=400&u=65aeee3d5216441f5aeda0e1daf92bb0ad8fae7a&v=4'; 
const githubLink = 'https://github.com/dech53';

const router = useRouter();

const logout = () => {
  localStorage.removeItem('authToken');
  router.push('/login');
};
</script>

<style scoped>
.header {
  width: 100%;
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 10px 20px;
  background-color: #282c34;
  color: white;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.2);
  border-radius: 8px;
}

.left-section {
  display: flex;
  align-items: center;
}

.avatar {
  width: 50px;
  height: 50px;
  border-radius: 50%;
  margin-right: 15px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.2);
}

.github-button {
  background-color: #87c0ca;
  color: white;
  padding: 8px 16px;
  text-decoration: none;
  border-radius: 5px;
  transition: background-color 0.3s;
}

.github-button:hover {
  background-color: #108b96;
}

.logout-button {
  background-color: #ff4d4f;
  color: white;
  padding: 8px 16px;
  border: none;
  border-radius: 5px;
  cursor: pointer;
  transition: background-color 0.3s;
}

.logout-button:hover {
  background-color: #d9363e;
}
</style>
