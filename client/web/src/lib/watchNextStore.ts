import { writable, derived } from 'svelte/store';

type WatchStatus = {
	watchStatusId: string;
	relationId: string;
	relationType: string;
};

export const apiData = writable<WatchStatus[]>([]);

export const watchNext = derived(apiData, ($apiData) => {
	return $apiData;
});
