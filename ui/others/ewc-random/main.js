document.addEventListener('DOMContentLoaded', () => {
    const videos = [
        '/ui/others/ewc-random/1.mp4',
        '/ui/others/ewc-random/2.mp4',
        '/ui/others/ewc-random/3.mp4',
        '/ui/others/ewc-random/4.mp4',
        '/ui/others/ewc-random/5.mp4',
        '/ui/others/ewc-random/6.mp4',
        '/ui/others/ewc-random/7.mp4',
        '/ui/others/ewc-random/8.mp4',
    ];

    const randomIndex = Math.floor(Math.random() * videos.length);
    const selectedVideo = videos[randomIndex];

    const videoElement = document.createElement('video');
    videoElement.src = selectedVideo;
    videoElement.autoplay = true;
    videoElement.controls = false;
    videoElement.style.position = 'fixed';
    videoElement.style.top = '0';
    videoElement.style.left = '0';
    videoElement.style.width = '100%';
    videoElement.style.height = '100%';
    videoElement.style.zIndex = '9999';
    videoElement.style.objectFit = 'cover';
    
    document.body.appendChild(videoElement);
    
    videoElement.requestFullscreen().catch(err => {
        console.error(`Error attempting to enable full-screen mode: ${err.message} (${err.name})`);
    });
});