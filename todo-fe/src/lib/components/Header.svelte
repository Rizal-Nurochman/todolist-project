<script>
    let isLogin=false;
    const navbar=[
        {
            id:1,
            value:"Home"
        },
        {
            id:2,
            value:"About"
        }
    ]

</script>

<header class="sticky top-0 left-0 w-full py-4 bg-gray-950 flex justify-between px-12 items-center">
    <h1 class="text-3xl font-bold text-blue-600">
        Slizer
    </h1>
    <nav class="flex items-center gap-4">

    </nav>
    {#if isLogin}
         <!-- content here -->
    {:else}
        <div class="flex items-cnter gap-4">
            <a href="/signin" class="px-5 py-2 bg-blue-600 text-white font-medium rounded-md hover:bg-blue-700 transition-all shadow-lg shadow-blue-600/20">
                Sign In
            </a>
            <a href="/signup" class="px-5 py-2 border border-blue-500 text-blue-500 font-medium rounded-md hover:bg-blue-600/10 transition-all">
                Sign Up
            </a>
        </div>
    {/if}
</header>