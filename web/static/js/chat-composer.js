(function () {
  function resizeComposer(textarea) {
    var maxHeight = Number(textarea.dataset.maxHeight || 220);
    textarea.style.height = "auto";
    var nextHeight = Math.min(textarea.scrollHeight, maxHeight);
    textarea.style.height = nextHeight + "px";
    textarea.style.overflowY = textarea.scrollHeight > maxHeight ? "auto" : "hidden";
    textarea.scrollTop = textarea.scrollHeight;
  }

  function bindComposer(textarea) {
    resizeComposer(textarea);
    textarea.addEventListener("input", function () {
      resizeComposer(textarea);
    });
  }

  function init() {
    document.querySelectorAll("[data-buildy-chat]").forEach(bindComposer);
  }

  if (document.readyState === "loading") {
    document.addEventListener("DOMContentLoaded", init);
  } else {
    init();
  }
})();
