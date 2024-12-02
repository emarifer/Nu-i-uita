<script lang="ts">
    import Router from "svelte-spa-router";
    // import { onMount } from "svelte";
    // import { locale, _ } from "svelte-i18n";
    // import { GetLanguage } from "../wailsjs/go/main/App";
    import Home from "./pages/Home.svelte";
    import Login from "./pages/Login.svelte";
    import About from "./pages/About.svelte";
    import AddPassword from "./pages/AddPassword.svelte";
    import Details from "./pages/Details.svelte";
    import EditPassword from "./pages/EditPassword.svelte";
    import Settings from "./pages/Settings.svelte";
    import { EventsEmit } from "../wailsjs/runtime/runtime";

    const routes = {
        "/": Login,
        "/home": Home,
        "/about": About,
        "/add": AddPassword,
        "/details/:id": Details,
        "/edit/:id": EditPassword,
        "/settings": Settings,
    };

    // onMount(() => {
    //     GetLanguage().then((result) => {
    //         locale.set(result);
    //         EventsEmit("change_title", `${$_("app_title")}`);
    //     });
    // });

    const onKeyDown = (e: KeyboardEvent) =>
        e.key === "Escape" ? EventsEmit("quit") : false;
</script>

<main class="overflow-hidden">
    <Router {routes} />
</main>

<svelte:window on:keydown={onKeyDown} />

<!-- 
    Toggle Password:
    https://svelte.dev/playground/3632c24a5c094528b961dc2c7ed119a4?version=3.32.3
    https://preline.co/docs/toggle-password.html

    Daisyui colors pallette & OKLCH Color Picker & Converter:
    https://unpkg.com/browse/daisyui@4.12.2/dist/themes.css
    https://oklch.com/
    https://github.com/saadeghi/daisyui/discussions/3096
 -->
