<script lang="ts">
  import AuthForm from "$lib/components/AuthForm.svelte";


</script>

<main class="w-full flex justify-center items-center py-8">
    <AuthForm
        msg="Signup"
        action="?/signup"
        defaultvalues={{
            email:'',
            password:'',
            confirmPassword:''
        }}
    />
</main>
