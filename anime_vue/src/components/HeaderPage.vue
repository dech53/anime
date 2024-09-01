<template>
  <header class="header">
    <div class="left-section">
      <img class="avatar" :src="avatarUrl" alt="Avatar" />
      <a :href="githubLink" class="link-button">
        <img src="https://cdn.icon-icons.com/icons2/2406/PNG/512/github_git_icon_145985.png" alt="GitHub" class="link-icon" />
      </a>
      <a :href="telegramLink" class="link-button">
        <img src="https://cdn.icon-icons.com/icons2/2429/PNG/512/telegram_logo_icon_147228.png" alt="Telegram" class="link-icon" />
      </a>
    </div>
      <button class="logout-button" @click="logout">登出</button>
  </header>
</template>

<script setup>
import { useRouter } from 'vue-router';
const avatarUrl = 'https://avatars.githubusercontent.com/u/93471086?s=400&u=65aeee3d5216441f5aeda0e1daf92bb0ad8fae7a&v=4';
const githubLink = 'https://github.com/dech53';
const telegramLink = 'https://t.me/dech53';
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
  width: 100wh;
  justify-content: space-between;
  align-items: center;
  padding: 10px 20px;
  background-color: linear-gradient(135deg, #17ead9, #6078ea);
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

.link-icon {
  width: 30px;
  height: 30px;
  border-radius: 50%;
  margin-right: 15px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.2);
}

.link-button {
  display: flex;
  align-items: center;
}

.logout-button {
  background-color: #ff6163;
  color: white;
  margin-right: 3px;
  padding: 8px 16px;
  border: none;
  border-radius: 5px;
  cursor: pointer;
  transition: background-color 0.3s;
}

.anime-add-button {
  background-color: #a9be7b;
  color: white;
  margin-right: 3px;
  padding: 8px 16px;
  border: none;
  border-radius: 5px;
  cursor: pointer;
  transition: background-color 0.3s;
}

.anime-add-button:hover {
  background-color: #2a6e3f;
}

.logout-button:hover {
  background-color: #d9363e;
}
</style>
