<script lang="ts">
    import { _ } from "svelte-i18n";
    import { DeleteEntry } from "../../wailsjs/go/main/App";
    import type { models } from "../../wailsjs/go/models";
    import { push } from "svelte-spa-router";
    import { showWarning } from "./popups/popups";

    let { entry }: { entry: models.PasswordEntry } = $props();

    const showAlert = (website: string, id: string) => {
        let data: string[] = [
            `${$_("alert_deleting_password")} "${website}."`,
            `${$_("alert_confirm_deleting")}"`,
        ];
        showWarning(data).then((result) => {
            if (result.value) {
                DeleteEntry(id).then(() => push("/home"));
            }
        });
    };
</script>

<div class="flex gap-2 absolute bottom-2 right-6">
    <button
        data-tip={$_("edit_btn_tip")}
        class="btn btn-ghost btn-sm tooltip tooltip-top"
        aria-label="Edit Item"
        onclick={() => push(`/edit/${entry.Id}`)}
    >
        <svg
            xmlns="http://www.w3.org/2000/svg"
            fill="#7480ff"
            class="w-4"
            viewBox="0 0 16 16"
        >
            <path
                d="M12.854.146a.5.5 0 0 0-.707 0L10.5 1.793 14.207 5.5l1.647-1.646a.5.5 0 0 0 0-.708zm.646 6.061L9.793 2.5 3.293 9H3.5a.5.5 0 0 1 .5.5v.5h.5a.5.5 0 0 1 .5.5v.5h.5a.5.5 0 0 1 .5.5v.5h.5a.5.5 0 0 1 .5.5v.207zm-7.468 7.468A.5.5 0 0 1 6 13.5V13h-.5a.5.5 0 0 1-.5-.5V12h-.5a.5.5 0 0 1-.5-.5V11h-.5a.5.5 0 0 1-.5-.5V10h-.5a.5.5 0 0 1-.175-.032l-.179.178a.5.5 0 0 0-.11.168l-2 5a.5.5 0 0 0 .65.65l5-2a.5.5 0 0 0 .168-.11z"
            />
        </svg>
    </button>

    <button
        data-tip={$_("delete_btn_tip")}
        class="btn btn-ghost btn-sm tooltip tooltip-top"
        aria-label="Delete Item"
        onclick={() => showAlert(entry.Website, entry.Id)}
    >
        <svg
            fill="#7480ff"
            class="w-4"
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 16 16"
        >
            <path
                d="M2.5 1a1 1 0 0 0-1 1v1a1 1 0 0 0 1 1H3v9a2 2 0 0 0 2 2h6a2 2 0 0 0 2-2V4h.5a1 1 0 0 0 1-1V2a1 1 0 0 0-1-1H10a1 1 0 0 0-1-1H7a1 1 0 0 0-1 1zm3 4a.5.5 0 0 1 .5.5v7a.5.5 0 0 1-1 0v-7a.5.5 0 0 1 .5-.5M8 5a.5.5 0 0 1 .5.5v7a.5.5 0 0 1-1 0v-7A.5.5 0 0 1 8 5m3 .5v7a.5.5 0 0 1-1 0v-7a.5.5 0 0 1 1 0"
            />
        </svg>
    </button>
</div>
