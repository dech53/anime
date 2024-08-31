<template>
  <div id="app">
    <br>
    <div v-for="anime in animes" :key="anime.Name" class="anime-item">
      <img :src="anime.Cover" :alt="anime.Name" class="cover" @click="selectAnime(anime)" />
      <div>
        <h2 @click="selectAnime(anime)">
          {{ truncateName(anime.Name) }}
        </h2>
        <p>Season: {{ anime.Season }}</p>
        <p>
          Episodes:
          <span>{{ anime.EpisodeCount }}</span>
        </p>
      </div>
    </div>
    <!-- <video id="video-player" controls autoplay></video> -->
  </div>
</template>

<script>

export default {
  data() {
    return {
      animes: [],
    };
  },
  methods: {
    async fetchAnimeInfo() {
      try {
        const token = localStorage.getItem('authToken');
        const response = await this.$axios.get('http://localhost:1226/info/getInfos', {
          headers: {
            'Authorization': token
          }
        });
        if (response.data.code === 302) {
          // 处理 302 重定向
          this.$router.push('/login'); // 重定向到登录页面
        } else {
          this.animes = response.data.result;
        }
      } catch (error) {
        console.error('Failed to fetch anime info:', error);
      }
    },
    selectAnime(anime) {
      console.log(anime);
    },
    handleStorageChange(event) {
      if (event.key === 'authToken' && event.newValue === null) {
        // 如果 authToken 被删除，重定向到登录页面
        this.$router.push('/login');
      }
    },
    truncateName(name) {
      return name.length > 8 ? name.slice(0, 8) + '..' : name;
    },
  },
  mounted() {
    this.fetchAnimeInfo();
    window.addEventListener('storage', this.handleStorageChange);
  },
};
</script>

<style>
#app {
  display: flex;
  flex-direction: row;
  align-items: center;
  flex-wrap: wrap;
}

body {
  font-family: Arial, sans-serif;
  background-color: #f4f4f4;
  margin: 0;
  padding: 20px;
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
}

h1 {
  color: #333;
  margin-bottom: 20px;
  text-align: center;
}

.anime-item {
  display: flex;
  align-items: flex-start;
  flex-direction: column;
  background-color: #fff;
  border-radius: 8px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  margin-bottom: 20px;
  margin: 5px;
  padding: 15px;
  width: 250px;
  height: 245px;
  transition: transform 0.2s;
}

.anime-item:hover {
  transform: scale(1.02);
}

.cover {
  width: 120px;
  height: 180px;
  margin-right: 20px;
  cursor: pointer;
  border-radius: 8px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

h2 {
  cursor: pointer;
  margin: 0;
  font-size: 24px;
  color: #333;
}

p {
  margin: 5px 0;
  color: #666;
}

.episode-number {
  margin-right: 10px;
  cursor: pointer;
  background-color: #007bff;
  color: #fff;
  padding: 5px 10px;
  border-radius: 4px;
  transition: background-color 0.2s;
}

.episode-number:hover {
  background-color: #0056b3;
}

video {
  width: 60%;
  height: auto;
  border-radius: 8px;
  margin-top: 30px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.2);
}

@media (max-width: 768px) {
  .anime-item {
    flex-direction: column;
    align-items: flex-start;
  }

  .cover {
    width: 100px;
    height: 150px;
    margin-bottom: 10px;
  }

  video {
    width: 100%;
  }
}
</style>
