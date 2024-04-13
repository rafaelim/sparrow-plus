import { writable, derived } from 'svelte/store';

type Movies = {
	movieId: string;
	name: string;
	year: string;
	path: string;
};

export const apiData = writable<Movies[]>([]);

export const movies = derived(apiData, ($apiData) => {
	return $apiData;
});
