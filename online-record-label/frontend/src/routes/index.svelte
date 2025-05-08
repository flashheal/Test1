<script>
  export let title = "Welcome to Our Online Record Label";
</script>

<main class="flex flex-col items-center justify-center min-h-screen bg-gray-100">
  <h1 class="text-4xl font-bold text-center text-blue-600">{title}</h1>
  <p class="mt-4 text-lg text-gray-700">
    Discover new music, artists, and albums. Join us in celebrating the art of sound!
  </p>
  <div class="mt-8">
    <a href="/artists" class="px-4 py-2 text-white bg-blue-500 rounded hover:bg-blue-600">
      Explore Artists
    </a>
    <a href="/albums" class="ml-4 px-4 py-2 text-white bg-green-500 rounded hover:bg-green-600">
      View Albums
    </a>
  </div>
</main>