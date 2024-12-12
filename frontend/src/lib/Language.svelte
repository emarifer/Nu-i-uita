<script lang="ts">
    import { onMount } from "svelte";
    import { locale, _ } from "svelte-i18n";
    import { SaveLanguage } from "../../wailsjs/go/main/App";
    import { EventsEmit } from "../../wailsjs/runtime/runtime";

    type language = {
        code: string;
        name: string;
    };

    const languages: language[] = [
        { code: "en", name: "English" },
        { code: "es", name: "Español" },
    ];

    let { languageName }: { languageName: string } = $props();

    onMount(() => (selected.name = languageName));

    let selected = $state({ code: "", name: "" } as language);
    const handleChange = (newLocale: language) => {
        // console.log("Language selected:", newLocale);
        locale.set(newLocale["code"]);
        EventsEmit(
            "change_titles",
            $_("app_title"),
            $_("select_directory"),
            $_("select_file"),
        );
        SaveLanguage(newLocale["code"]);
    };
</script>

<select
    bind:value={selected}
    onchange={() => handleChange(selected)}
    class="absolute top-2 right-2 select select-sm select-primary w-fit max-w-xs"
>
    {#if selected.code == ""}
        <option disabled value={selected}>{languageName}</option>
    {:else}
        <option disabled>🌐 {$_("language")}</option>
    {/if}
    {#each languages as language}
        <option value={language}>{language.name}</option>
    {/each}
</select>
