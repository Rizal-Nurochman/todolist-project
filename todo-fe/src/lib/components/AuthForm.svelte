<script lang="ts">
    import type { FormProps } from '$lib/types';

    let { msg, defaultvalues, action }: FormProps = $props();
</script>

<form action={action} method="post" class="space-y-6 p-8 bg-gray-900 rounded-lg shadow-xl border border-gray-800 w-full max-w-lg mx-auto">
    <div class="text-center mb-4">
        <h2 class="text-2xl font-bold text-white">{msg}</h2>
        <p class="text-gray-400 text-sm mt-1">Please enter your credentials</p>
    </div>
    
    <div class="space-y-4">
        <div class="flex flex-col gap-2">
            <label for="email" class="block text-sm font-medium text-gray-300 mb-1">Email Address</label>
            <div class="relative">
                <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                    <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 text-gray-500" viewBox="0 0 20 20" fill="currentColor">
                        <path d="M2.003 5.884L10 9.882l7.997-3.998A2 2 0 0016 4H4a2 2 0 00-1.997 1.884z" />
                        <path d="M18 8.118l-8 4-8-4V14a2 2 0 002 2h12a2 2 0 002-2V8.118z" />
                    </svg>
                </div>
                <input 
                    id="email" 
                    name="email" 
                    type="email" 
                    placeholder="you@example.com" 
                    value={defaultvalues?.email || ''}
                    required 
                    class="block w-full pl-10 pr-3 py-3 bg-gray-800 text-white placeholder-gray-500 border border-gray-700 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition duration-150"
                />
            </div>
        </div>
        
        <div class="flex flex-col gap-2">
            <label for="password" class="block text-sm font-medium text-gray-300 mb-1">Password</label>
            <div class="relative">
                <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                    <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 text-gray-500" viewBox="0 0 20 20" fill="currentColor">
                        <path fill-rule="evenodd" d="M5 9V7a5 5 0 0110 0v2a2 2 0 012 2v5a2 2 0 01-2 2H5a2 2 0 01-2-2v-5a2 2 0 012-2zm8-2v2H7V7a3 3 0 016 0z" clip-rule="evenodd" />
                    </svg>
                </div>
                <input 
                    id="password" 
                    name="password" 
                    type="password" 
                    placeholder="••••••••" 
                    value={defaultvalues?.password || ''}
                    required 
                    class="block w-full pl-10 pr-3 py-3 bg-gray-800 text-white placeholder-gray-500 border border-gray-700 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition duration-150"
                />
            </div>
        </div>
        
        {#if defaultvalues.hasOwnProperty('confirmPassword')}
            <div class="flex flex-col gap-2">
                <label for="confirmPassword" class="block text-sm font-medium text-gray-300 mb-1">Confirm Password</label>
                <div class="relative">
                    <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 text-gray-500" viewBox="0 0 20 20" fill="currentColor">
                            <path fill-rule="evenodd" d="M2.166 4.999A11.954 11.954 0 0010 1.944 11.954 11.954 0 0017.834 5c.11.65.166 1.32.166 2.001 0 5.225-3.34 9.67-8 11.317C5.34 16.67 2 12.225 2 7c0-.682.057-1.35.166-2.001zm11.541 3.708a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd" />
                        </svg>
                    </div>
                    <input 
                        id="confirmPassword" 
                        name="confirmPassword" 
                        type="password" 
                        placeholder="••••••••" 
                        value={defaultvalues?.confirmPassword || ''}
                        required 
                        class="block w-full pl-10 pr-3 py-3 bg-gray-800 text-white placeholder-gray-500 border border-gray-700 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition duration-150"
                    />
                </div>
            </div>
        {/if}
    </div>
    
    {#if !(defaultvalues.hasOwnProperty('confirmPassword'))}
        <div class="flex items-center justify-between text-sm">
            <div class="flex items-center">
                <input id="remember_me" name="remember_me" type="checkbox" class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-700 rounded bg-gray-800">
                <label for="remember_me" class="ml-2 block text-gray-300">Remember me</label>
            </div>
            <a href="/forgot-password" class="text-blue-500 hover:text-blue-400 transition">Forgot password?</a>
        </div>
    {/if}
    
    <button 
        type="submit" 
        class="w-full flex justify-center py-3 px-4 border border-transparent rounded-md shadow-sm text-base font-medium text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-offset-gray-900 focus:ring-blue-500 transition duration-150 cursor-pointer"
    >
        {msg}
    </button>
    
    {#if !(defaultvalues.hasOwnProperty('confirmPassword'))}
        <div class="text-center text-sm text-gray-400">
            Don't have an account? <a href="/signup" class="text-blue-500 hover:text-blue-400 transition">Sign up</a>
        </div>
    {:else if msg.toLowerCase().includes('sign up')}
        <div class="text-center text-sm text-gray-400">
            Already have an account? <a href="/signin" class="text-blue-500 hover:text-blue-400 transition">Sign in</a>
        </div>
    {/if}
</form>