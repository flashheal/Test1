import sveltePreprocess from 'svelte-preprocess';

export default {
  preprocess: sveltePreprocess(),
  kit: {
    // Specify the adapter for deployment
    adapter: {
      name: '@sveltejs/adapter-auto',
      options: {
        // Options for the adapter
      }
    },
    // Other configurations can go here
  }
};