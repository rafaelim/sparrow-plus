<script lang="ts">
	import 'video.js/dist/video-js.css';
	import videoJS from 'video.js';

	let scrollbox: Element;
	$: if (scrollbox) {
		const player = videoJS(scrollbox, {
			controls: true,
			autoplay: true,
			aspectRatio: '16:9',
			html5: {
				nativeTextTracks: false,
				nativeAudioTracks: false,
				vhs: {
					overrideNative: true
				}
			},
			sources: [
				{
					src: 'http://localhost:3000/api/stream/video.mkv',
					type: 'application/x-mpegURL'
				}
			]
		});
	}
</script>

<video bind:this={scrollbox} controls={true} class="vjs-matrix video-js"> </video>
<div id="audioTrackControl"></div>
