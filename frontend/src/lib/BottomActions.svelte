<script lang="ts">
    import { push } from "svelte-spa-router";
    import { _ } from "svelte-i18n";
</script>

<div class="absolute bottom-2 left-2 flex gap-1">
    <button
        data-tip={$_("settings_btn")}
        class="btn btn-ghost btn-sm tooltip tooltip-top"
        aria-label="Settings"
        onclick={() => push("/settings")}
    >
        <svg
            xmlns="http://www.w3.org/2000/svg"
            class="w-5"
            fill="#7480ff"
            viewBox="0 0 16 16"
        >
            <path
                d="M9.405 1.05c-.413-1.4-2.397-1.4-2.81 0l-.1.34a1.464 1.464 0 0 1-2.105.872l-.31-.17c-1.283-.698-2.686.705-1.987 1.987l.169.311c.446.82.023 1.841-.872 2.105l-.34.1c-1.4.413-1.4 2.397 0 2.81l.34.1a1.464 1.464 0 0 1 .872 2.105l-.17.31c-.698 1.283.705 2.686 1.987 1.987l.311-.169a1.464 1.464 0 0 1 2.105.872l.1.34c.413 1.4 2.397 1.4 2.81 0l.1-.34a1.464 1.464 0 0 1 2.105-.872l.31.17c1.283.698 2.686-.705 1.987-1.987l-.169-.311a1.464 1.464 0 0 1 .872-2.105l.34-.1c1.4-.413 1.4-2.397 0-2.81l-.34-.1a1.464 1.464 0 0 1-.872-2.105l.17-.31c.698-1.283-.705-2.686-1.987-1.987l-.311.169a1.464 1.464 0 0 1-2.105-.872zM8 10.93a2.929 2.929 0 1 1 0-5.86 2.929 2.929 0 0 1 0 5.858z"
            />
        </svg>
    </button>
    <button
        data-tip={$_("about_btn")}
        class="btn btn-ghost btn-sm tooltip tooltip-top"
        aria-label="About"
        onclick={() => push("/about")}
    >
        <svg
            xmlns="http://www.w3.org/2000/svg"
            class="w-5"
            fill="#7480ff"
            viewBox="0 0 16 16"
        >
            <path
                d="M8 15A7 7 0 1 1 8 1a7 7 0 0 1 0 14m0 1A8 8 0 1 0 8 0a8 8 0 0 0 0 16"
            />
            <path
                d="m8.93 6.588-2.29.287-.082.38.45.083c.294.07.352.176.288.469l-.738 3.468c-.194.897.105 1.319.808 1.319.545 0 1.178-.252 1.465-.598l.088-.416c-.2.176-.492.246-.686.246-.275 0-.375-.193-.304-.533zM9 4.5a1 1 0 1 1-2 0 1 1 0 0 1 2 0"
            />
        </svg>
    </button>
</div>

<button
    data-tip={$_("add_btn")}
    class="absolute bottom-2 right-2 btn btn-ghost btn-sm tooltip tooltip-top flex gap-2 text-primary"
    onclick={() => push("/add")}
>
    <svg
        xmlns="http://www.w3.org/2000/svg"
        class="w-5"
        fill="#7480ff"
        viewBox="0 0 16 16"
    >
        <path
            fill-rule="evenodd"
            d="M8 2a.5.5 0 0 1 .5.5v5h5a.5.5 0 0 1 0 1h-5v5a.5.5 0 0 1-1 0v-5h-5a.5.5 0 0 1 0-1h5v-5A.5.5 0 0 1 8 2"
        />
    </svg>
    {$_("add_btn_text")}
</button>
