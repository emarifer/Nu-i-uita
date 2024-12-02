<script lang="ts">
    import { pop } from "svelte-spa-router";
    import { _ } from "svelte-i18n";
</script>

<button
    data-tip={$_("back")}
    class="absolute bottom-2 left-2 btn btn-ghost btn-sm tooltip tooltip-top flex gap-2 text-primary"
    aria-label="Back"
    onclick={() => pop()}
>
    <svg
        xmlns="http://www.w3.org/2000/svg"
        fill="#7480ff"
        class="w-5"
        viewBox="0 0 16 16"
    >
        <path
            fill-rule="evenodd"
            d="M15 8a.5.5 0 0 0-.5-.5H2.707l3.147-3.146a.5.5 0 1 0-.708-.708l-4 4a.5.5 0 0 0 0 .708l4 4a.5.5 0 0 0 .708-.708L2.707 8.5H14.5A.5.5 0 0 0 15 8"
        />
    </svg>
</button>
