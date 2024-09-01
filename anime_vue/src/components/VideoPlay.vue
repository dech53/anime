<template>
    <div class="main_menu">
        <div v-for="a in anime" :key="a.Name" class="anime-item">
            <img :src="a.Cover" :alt="a.Name" class="cover" @click="ChooseEpisode(a)" />
            <div>
                <h2>{{ a.Name }}</h2>
                <p>{{ a.Author }}</p>
                <p>Season: {{ a.Season }}</p>
                <p>集数:</p>
                <div class="episode-container">
                    <span v-for="index in a.EpisodeCount" :key="index" @click="playEpisode(a, index)"
                        class="EpisodeChoice">
                        {{ index }}
                    </span>
                </div>
            </div>
        </div>
        <video class="video_play" ref="videoPlayer" controls></video>
    </div>
</template>

<script>
import { useRoute } from 'vue-router';
import Hls from 'hls.js'; // 如果你使用 npm 安装 hls.js

export default {
    data() {
        return {
            anime: [],
            hls: null, // 添加 hls 实例的引用
        }
    },
    methods: {
        async fetchInfo() {
            try {
                const router = useRoute();
                const token = localStorage.getItem('authToken');
                const response = await this.$axios.get('http://localhost:1226/info/getInfoById/' + router.params.id, {
                    headers: {
                        'Authorization': token
                    }
                });
                if (response.data.code == 200) {
                    this.anime = response.data.result;
                    // 默认播放第一个集数
                    if (this.anime.length > 0 && this.anime[0].EpisodeCount > 0) {
                        this.playEpisode(this.anime[0], 1);
                    }
                }
            } catch (error) {
                console.log('Failed to fetch anime info:', error);
            }
        },

        async playEpisode(anime, episode) {
            try {
                const videoPath = `http://localhost:1226/hls/${anime.Year}/${anime.Name}/${anime.Season}/${episode}/output.m3u8`;
                console.log(videoPath);
                const videoPlayer = this.$refs.videoPlayer;
                // 清理之前的 HLS 实例
                if (this.hls) {
                    this.hls.destroy();
                }
                // 初始化 HLS.js
                if (Hls.isSupported()) {
                    this.hls = new Hls();
                    this.hls.loadSource(videoPath);
                    this.hls.attachMedia(videoPlayer);
                    this.hls.on(Hls.Events.MANIFEST_PARSED, () => {
                        videoPlayer.play();
                    });
                    this.hls.on(Hls.Events.ERROR, (event, data) => {
                        console.error('HLS.js error:', data.fatal);
                    });
                } else if (videoPlayer.canPlayType('application/vnd.apple.mpegurl')) {
                    videoPlayer.src = videoPath;
                    videoPlayer.addEventListener('loadedmetadata', () => {
                        videoPlayer.play();
                    });
                } else {
                    console.error('This browser does not support HLS.');
                }

            } catch (error) {
                console.error('Failed to play video:', error);
            }
        }
    },
    mounted() {
        this.fetchInfo();
    }
}
</script>

<style scoped>
.main_menu {
    display: flex;
    flex-direction: row;
    align-items: flex-start;
    justify-content: center;
    height: 80vh;
    width: 78vw;
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
    height: 100%;
    flex-wrap: wrap;
    transition: transform 0.2s;
}

.episode-container {
    display: flex;
    flex-direction: row;
    flex-wrap: wrap;
    cursor: pointer;
}

.EpisodeChoice {
    padding: 5px;
    width: 20px;
    display: flex;
    align-items: center;
    justify-content: center;
    background-color: #29b6f6;
    margin: 5px;
    border-radius: 15%;
    text-align: center;
}

.video_play {
    width: 100%;
    height: 100%;
}
</style>
