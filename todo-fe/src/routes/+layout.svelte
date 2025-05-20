<script lang="ts">
  import "../app.css";
  import Header from "$lib/components/Header.svelte";

  let { children }=$props()

</script>

<!-- app layout -->
<Header/>
<main class="w-full py-8 min-h-screen bg-slate-900">
    {@render children()}
</main>
