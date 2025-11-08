// Disable logging
console.log = (args) => null;

// Initialize Monaco Editor
require.config({ paths: { 'vs': 'https://cdnjs.cloudflare.com/ajax/libs/monaco-editor/0.34.0/min/vs' } });
require(['vs/editor/editor.main'], function() {
  // Define custom theme
  monaco.editor.defineTheme('modernTheme', {
    base: 'vs-dark',
    inherit: true,
    rules: [
      { token: 'keyword', foreground: '818cf8', fontStyle: 'bold' },
      { token: 'string', foreground: '34d399', fontStyle: 'italic' },
      { token: 'number', foreground: 'f472b6' },
      { token: 'comment', foreground: '64748b', fontStyle: 'italic' },
      { token: 'type', foreground: '38bdf8' },
      { token: 'delimiter', foreground: 'cbd5e1' },
      { token: 'delimiter.bracket', foreground: '94a3b8' },
    ],
    colors: {
      'editor.background': '#1e293b',
      'editor.foreground': '#f1f5f9',
      'editorCursor.foreground': '#f1f5f9',
      'editor.lineHighlightBackground': '#334155',
      'editorLineNumber.foreground': '#64748b',
      'editor.selectionBackground': '#475569',
      'editor.inactiveSelectionBackground': '#334155',
      'editorIndentGuide.background': '#334155',
      'editorIndentGuide.activeBackground': '#475569',
    }
  });

  // Sample code
  const sampleCodes = {
    user: `User(
      name=Alice, 
      age=30, 
      active=true
)`,
    product: `Product(
      id=102, 
      name=Laptop, 
      price=999.99, 
      inStock=true
)`,
    address: `Address(
      street=123 Elm St, 
      city=Springfield, 
      zip=12345
)`,
    complex: `Order(
    id=5001, 
    user=User(
        name=Alice, 
        age=30, 
        active=true
    ), 
    items=[
        Product(id=102, name=Laptop, price=999.99, inStock=true), 
        Product(id=205, name=Mouse, price=19.99, inStock=false)
    ], 
    total=1019.98
)`
  };

  // Initialize Java editor
  const javaEditor = monaco.editor.create(document.getElementById('javaEditor'), {
    value: sampleCodes.user,
    language: 'java',
    theme: 'modernTheme',
    automaticLayout: true,
    minimap: { enabled: false },
    scrollBeyondLastLine: false,
    fontSize: 14,
    folding: true,
    lineNumbers: 'on',
    renderIndentGuides: true,
    contextmenu: true,
    scrollbar: {
      useShadows: false,
      verticalScrollbarSize: 10,
      horizontalScrollbarSize: 10
    }
  });

  // Initialize JSON editor
  const jsonEditor = monaco.editor.create(document.getElementById('jsonEditor'), {
    value: '// JSON output will appear here',
    language: 'json',
    theme: 'modernTheme',
    automaticLayout: true,
    minimap: { enabled: false },
    scrollBeyondLastLine: false,
    fontSize: 14,
    readOnly: false,
    folding: true,
    lineNumbers: 'on',
    renderIndentGuides: true,
    contextmenu: true,
    autoIndent: true,
    formatOnPaste: true,
    scrollbar: {
      useShadows: false,
      verticalScrollbarSize: 10,
      horizontalScrollbarSize: 10
    }
  });

  // Scroll control for editors
  let javaEditorFocused = false;
  let jsonEditorFocused = false;

  // Track focus state for Java editor
  javaEditor.onDidFocusEditorWidget(() => {
    javaEditorFocused = true;
  });
  
  javaEditor.onDidBlurEditorWidget(() => {
    javaEditorFocused = false;
  });

  // Track focus state for JSON editor
  jsonEditor.onDidFocusEditorWidget(() => {
    jsonEditorFocused = true;
  });
  
  jsonEditor.onDidBlurEditorWidget(() => {
    jsonEditorFocused = false;
  });

  // More aggressive scroll control - capture at document level
  document.addEventListener('wheel', function(event) {
    // Check if the scroll is happening over an editor
    const target = event.target;
    const javaEditorContainer = document.getElementById('javaEditor');
    const jsonEditorContainer = document.getElementById('jsonEditor');
    
    const isOverJavaEditor = javaEditorContainer.contains(target);
    const isOverJsonEditor = jsonEditorContainer.contains(target);
    
    if (isOverJavaEditor && !javaEditorFocused) {
      // Prevent editor scroll, allow page scroll
      event.preventDefault();
      window.scrollBy(0, event.deltaY);
    } else if (isOverJsonEditor && !jsonEditorFocused) {
      // Prevent editor scroll, allow page scroll
      event.preventDefault();
      window.scrollBy(0, event.deltaY);
    }
  }, { passive: false, capture: true });

  // Function to convert Java/Lombok to JSON
  function convertToJson() {
    const lombokInput = javaEditor.getValue();
    console.log(lombokInput);
    let jsonOutput = lombokToJson(lombokInput);
    jsonEditor.setValue(jsonOutput);

    // edit mode
    jsonEditor.updateOptions({ readOnly: false });

    jsonEditor
      .getAction('editor.action.formatDocument')
      .run()
      .then(() => {
        // read only
        jsonEditor.updateOptions({ readOnly: true })
        console.log("Set it to readonly again");
      });
  }

  function beautifyInput() {
    if (typeof beautifyLombok !== 'function') {
      console.warn('Beautify function is not ready yet.');
      return;
    }

    const lombokInput = javaEditor.getValue();
    const beautified = beautifyLombok(lombokInput, 2);
    if (beautified) {
      javaEditor.setValue(beautified);
    }
  }

  // Sample dropdown toggle
  const examplesDropdown = document.getElementById('examples-dropdown');
  const samplesMenu = document.getElementById('samples-menu');

  examplesDropdown.addEventListener('click', function(e) {
    e.stopPropagation();
    samplesMenu.classList.toggle('show');
  });

  // Close dropdown when clicking elsewhere
  document.addEventListener('click', function(event) {
    if (!examplesDropdown.contains(event.target)) {
      samplesMenu.classList.remove('show');
    }
  });

  // Sample selection
  document.getElementById('userSample').addEventListener('click', function() {
    javaEditor.setValue(sampleCodes.user);
    samplesMenu.classList.remove('show');
  });

  document.getElementById('productSample').addEventListener('click', function() {
    javaEditor.setValue(sampleCodes.product);
    samplesMenu.classList.remove('show');
  });

  document.getElementById('addressSample').addEventListener('click', function() {
    javaEditor.setValue(sampleCodes.address);
    samplesMenu.classList.remove('show');
  });

  document.getElementById('complexSample').addEventListener('click', function() {
    javaEditor.setValue(sampleCodes.complex);
    samplesMenu.classList.remove('show');
  });

  // Convert button
  document.getElementById('convertButton').addEventListener('click', convertToJson);
  document.getElementById('beautifyButton').addEventListener('click', beautifyInput);

  // Copy button
  document.getElementById('copyButton').addEventListener('click', function() {
    const jsonContent = jsonEditor.getValue();
    navigator.clipboard.writeText(jsonContent).then(function() {
      // Show success message
      const copyButton = document.getElementById('copyButton');
      const originalText = copyButton.innerHTML;
      copyButton.innerHTML = `
                        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
                        </svg>
                        <span>Copied!</span>
                    `;

      setTimeout(function() {
        copyButton.innerHTML = originalText;
      }, 2000);
    });
  });

  // Initial conversion
  setTimeout(convertToJson, 500);
});
