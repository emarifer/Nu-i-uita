
import Swal, { type SweetAlertResult } from "sweetalert2";
import { ERROR, QUESTION, SUCCESS, WARNING } from "./alert-icons";

export const showWarning: (data: string[]) => Promise<SweetAlertResult<any>> = (data) => Swal.fire({
    title: data[0],
    width: 275,
    icon: "warning",
    iconHtml: WARNING,
    background: "#1D232A",
    color: "#A6ADBA",
    showCancelButton: true,
    confirmButtonColor: "#3085d6",
    cancelButtonColor: "#d33",
    confirmButtonText: data[1],
    customClass: {
        title: "alertTitle",
        confirmButton: "alertConfirm",
        cancelButton: "alertCancel",
        icon: "alertIcon",
    },
});

export const showSuccess: (data: string) => Promise<SweetAlertResult<any>> = (data) => Swal.fire({
    title: data,
    width: 275,
    icon: "success",
    iconHtml: SUCCESS,
    background: "#1D232A",
    color: "#A6ADBA",
    confirmButtonColor: "#3085d6",
    customClass: {
        title: "alertTitle",
        confirmButton: "alertConfirm",
        icon: "alertIcon",
    },
});

export const showError: (data: string) => Promise<SweetAlertResult<any>> = (data) => Swal.fire({
    title: data,
    width: 275,
    icon: "error",
    iconHtml: ERROR,
    background: "#1D232A",
    color: "#A6ADBA",
    confirmButtonColor: "#3085d6",
    customClass: {
        title: "errorTitle",
        confirmButton: "alertConfirm",
        icon: "alertIcon",
    },
});

export const showQuestion: (data: string) => Promise<SweetAlertResult<any>> = (data) => Swal.fire({
    title: data,
    width: 275,
    input: "password",
    icon: "question",
    iconHtml: QUESTION,
    background: "#1D232A",
    color: "#A6ADBA",
    confirmButtonColor: "#3085d6",
    customClass: {
        title: "alertTitle",
        confirmButton: "alertConfirm",
        icon: "alertIcon",
        input: "inputDialog",
    },
});

