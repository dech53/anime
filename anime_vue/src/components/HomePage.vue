<template>
  <div id="app">
    <br>
    <div v-for="anime in animes" :key="anime.Name" class="anime-item">
      <img :src="anime.Cover" :alt="anime.Name" class="cover" @click="selectAnime(anime)" />
      <div>
        <h2 @click="selectAnime(anime)">
          {{ truncateName(anime.Name) }}
        </h2>
        <p>{{ anime.Season }}</p>
        <p>
          集数:
          <span>{{ anime.EpisodeCount }}</span>
        </p>
      </div>
    </div>
    <Modal :visible="showModal" :anime="newAnime" @close="showModal = false" @submit="addAnime" />
    <button class="test" @click="showModal = true">+</button>
  </div>
</template>

<script>
import Modal from './AddAnimeForm.vue';

export default {
  components: {
    Modal
  },
  data() {
    return {
      animes: [],
      showModal: false,
      newAnime: {
        Year: '',
        Name: '',
        Season: '',
        EpisodeCount: '',
        Author: '',
        Path: '',
        Cover: ''
      }
    };
  },
  methods: {
    async fetchAnimeInfo() {
      try {
        const token = localStorage.getItem('authToken');
        const response = await this.$axios.get('http://192.168.1.4:1226/info/getInfos', {
          headers: {
            'Authorization': token
          }
        });
        if (response.data.code === 302) {
          this.$router.push('/login');
        } else {
          this.animes = response.data.result;
        }
      } catch (error) {
        console.error('Failed to fetch anime info:', error);
      }
    },
    selectAnime(anime) {
      this.$router.push('/videoplay/' + anime.id);
    },
    handleStorageChange(event) {
      if (event.key === 'authToken' && event.newValue === null) {
        this.$router.push('/login');
      }
    },
    truncateName(name) {
      return name.length > 8 ? name.slice(0, 8) + '..' : name;
    },
    addAnime(anime) {
      if (anime.Name && anime.Year && anime.Season) {
        this.animes.push(anime);
        this.newAnime = {
          Year: '',
          Name: '',
          Season: '',
          EpisodeCount: '',
          Author: '',
          Path: '',
          Cover: ''
        };
      }
    }
  },
  mounted() {
    this.fetchAnimeInfo();
    window.addEventListener('storage', this.handleStorageChange);
  },
};
</script>

<style>
.modal {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 9999;
}

.modal-content {
  background: white;
  padding: 20px;
  border-radius: 8px;
  width: 80%;
  max-width: 600px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.modal-content h2 {
  margin-top: 0;
}

.modal-content form {
  display: flex;
  flex-direction: column;
}

.modal-content label {
  margin-bottom: 10px;
}

.modal-content button {
  margin-top: 10px;
  padding: 10px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

.modal-content button[type="submit"] {
  background-color: #007bff;
  color: white;
}

.modal-content button[type="button"] {
  background-color: #ccc;
  color: black;
}

#app {
  display: flex;
  flex-direction: row;
  align-items: center;
  flex-wrap: wrap;
  z-index: 1;
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

.test {
  position: fixed;
  bottom: 20px;
  right: 20px;
  background-color: #007bff;
  color: white;
  width: 60px;
  height: 60px;
  border: none;
  border-radius: 50%;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: background-color 0.3s, transform 0.2s;
  font-size: 24px;
  text-align: center;
  z-index: 5;
}

.test:hover {
  background-color: #0056b3;
  transform: scale(1.1);
}
</style>
