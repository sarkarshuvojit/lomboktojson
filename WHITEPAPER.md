# LombokToJson: A Technical Whitepaper

## Abstract

LombokToJson is a lightweight utility that converts Lombok's default `toString()` output format to standard JSON format. This whitepaper presents a comprehensive technical analysis of the system's architecture, focusing on two critical components: the custom tokenizer implementation and the innovative browser integration using WebAssembly (WASM) via Go's native compilation capabilities.

The tokenizer employs a single-pass, context-sensitive scanning algorithm that efficiently processes Lombok's object notation syntax. The browser integration leverages Go's built-in WebAssembly compilation to deliver the same parsing logic across multiple platforms without code duplication.

---

## 1. Tokenizer Architecture and Implementation

### 1.1 Overview

The LombokToJson tokenizer implements a custom lexical analyzer designed specifically for parsing Lombok's `toString()` output format. Unlike general-purpose JSON parsers, this tokenizer is optimized for the specific characteristics of Lombok's object notation, which includes class names, nested structures, and array literals.

### 1.2 Token Type System

The tokenizer defines a comprehensive set of 10 token types that capture the semantic elements of Lombok notation:

#### Structural Delimiters
- **PAREN_OPEN** (`(`) - Opening parenthesis for class constructors
- **PAREN_CLOSE** (`)`) - Closing parenthesis for class constructors  
- **ARRAY_OPEN** (`[`) - Opening bracket for array literals
- **ARRAY_CLOSE** (`]`) - Closing bracket for array literals
- **COMMA** (`,`) - Field separator in parameter lists
- **EQUALS** (`=`) - Assignment operator between keys and values

#### Semantic Content Types
- **CLASS_NAME** - Class/constructor names (identified by position before opening parenthesis)
- **KEY** - Parameter names (identified by position before equals sign)
- **VALUE** - Parameter values (identified by position after equals sign)
- **STRING_LITERAL** - Generic string content (fallback classification)
- **NUM_LITERAL** - Numeric literal values
- **EOF** - End-of-file marker

### 1.3 Scanner Algorithm

The scanner implements a **single-pass, character-by-character streaming algorithm** with sophisticated context analysis:

```go
type Scanner struct {
    sourceBytes    []byte    // Input byte stream
    curline        int       // Current line number tracking
    start          int       // Token start position
    end            int       // Current parsing position
    parenOpen      int       // Parenthesis depth counter
    
    // Literal accumulation state
    literalStarted bool
    literalStart   int
    literalEnd     int
    
    tokens         []types.Token  // Output token collection
}
```

#### Core Processing Logic

1. **Character Classification**: Each character is classified as either:
   - Structural delimiter requiring immediate token generation
   - Literal content to be accumulated into spans
   - Whitespace to be ignored

2. **Literal Accumulation**: The scanner accumulates consecutive literal characters:
   ```go
   func isLiteral(ch string) bool {
       return unicode.IsLetter([]rune(ch)[0]) || unicode.IsDigit([]rune(ch)[0])
   }
   ```

3. **Context-Sensitive Classification**: When a literal span terminates, the scanner performs lookahead and lookbehind analysis:
   ```go
   func (s *Scanner) stringLiteralToToken(literal string) types.Token {
       if s.sourceBytes[s.literalEnd+1] == '(' {  // Lookahead for class name
           return types.NewToken(types.CLASS_NAME, literal, nil, s.curline)
       }
       if s.sourceBytes[s.literalEnd+1] == '=' {  // Lookahead for key
           return types.NewToken(types.KEY, literal, nil, s.curline)
       }
       if s.sourceBytes[s.literalStart-1] == '=' { // Lookbehind for value
           return types.NewToken(types.VALUE, literal, nil, s.curline)
       }
       return types.NewToken(types.STRING_LITERAL, literal, nil, s.curline)
   }
   ```

### 1.4 Advanced Features

#### Nested Structure Support
The scanner maintains a `parenOpen` counter to track nesting depth, enabling proper parsing of complex nested structures:
```
Customer(addresses=[Address(street=123 Main St, city=Bangalore)])
```

#### Array Processing
Arrays are treated as first-class structural elements with dedicated token types for brackets, while array contents are processed using the same context-sensitive logic as other values.

#### Whitespace Handling
The implementation provides **implicit whitespace tolerance** by focusing only on literal and delimiter characters, enabling robust parsing of formatted input with arbitrary spacing.

### 1.5 Performance Characteristics

- **Time Complexity**: O(n) single-pass linear scanning
- **Space Complexity**: O(t) where t is the number of tokens generated
- **Memory Efficiency**: Character-by-character processing minimizes memory allocation
- **Error Recovery**: Graceful handling of malformed input through fallback classification

---

## 2. Browser Integration via WebAssembly (WASI)

### 2.1 Architecture Overview

LombokToJson employs Go's native WebAssembly compilation capabilities to deliver the same parsing logic in web browsers without requiring JavaScript reimplementation. This approach ensures **algorithmic consistency** across CLI, library, and web interface deployments.

### 2.2 WebAssembly Compilation Process

#### Build Process
The Go toolchain provides native WebAssembly compilation through the `js/wasm` target:

```bash
GOOS=js GOARCH=wasm go build -o lombok2json.wasm ./examples/lomboktojsonwasm/
```

This generates a self-contained WebAssembly module that includes:
- The complete tokenizer and parser implementation
- Go runtime environment
- Memory management system
- JavaScript interop layer

#### WASM Module Structure
The WebAssembly module exposes the core conversion functionality through a global JavaScript function:
```go
// examples/lomboktojsonwasm/main.go
func lombokToJson(lombokString string) string {
    result, err := l2j.LombokToJson(lombokString)
    if err != nil {
        return fmt.Sprintf("Error: %v", err)
    }
    return result
}

func main() {
    js.Global().Set("lombokToJson", js.FuncOf(lombokToJsonWrapper))
    select {} // Keep the program running
}
```

### 2.3 Browser Integration Implementation

#### JavaScript Runtime Integration
The web interface uses Go's provided `wasm_exec.js` runtime to instantiate and manage the WebAssembly module:

```javascript
const go = new Go();
WebAssembly
    .instantiateStreaming(fetch("assets/lombok2json.wasm"), go.importObject)
    .then(result => {
        go.run(result.instance);
    });
```

#### Function Invocation
Once initialized, the web interface can directly invoke the Go-compiled function:
```javascript
function convertToJson() {
    const lombokInput = javaEditor.getValue();
    let jsonOutput = lombokToJson(lombokInput);  // Direct WASM function call
    jsonEditor.setValue(jsonOutput);
}
```

### 2.4 Deployment Strategy

#### Manual Asset Management
The current deployment strategy employs a **manual copy process** that ensures explicit control over the WebAssembly module version:

1. **Compilation**: Go builds the WASM module in the development environment
2. **Manual Copy**: The generated `.wasm` file is manually copied to `docs/assets/`
3. **Version Control**: Both source and compiled assets are tracked in the repository
4. **Deployment**: Static hosting serves the WASM module alongside other web assets

#### Benefits of Manual Deployment
- **Explicit Control**: Manual copying ensures intentional updates to the web interface
- **Version Consistency**: Prevents automatic deployment of untested WASM builds
- **Debugging**: Easier to track which version of the parser is deployed
- **Simplicity**: No complex build pipeline or CI/CD integration required

### 2.5 Performance Characteristics

#### Startup Performance
- **Module Size**: Compact WASM module (~2-3MB including Go runtime)
- **Initialization Time**: Near-instantaneous loading with WebAssembly streaming compilation
- **First Parse**: Minimal cold-start overhead due to pre-compiled nature

#### Runtime Performance
- **Parsing Speed**: Near-native performance due to WebAssembly's efficient execution model
- **Memory Usage**: Efficient memory management through Go's garbage collector
- **Consistency**: Identical performance characteristics to native Go implementation

#### Browser Compatibility
- **Modern Browser Support**: Compatible with all browsers supporting WebAssembly (95%+ coverage)
- **Progressive Enhancement**: Graceful degradation possible with JavaScript fallback
- **Mobile Performance**: Efficient execution on mobile browsers due to WASM optimization

### 2.6 Advantages of the WASI Approach

#### Code Reuse
- **Single Implementation**: The same Go codebase serves CLI, library, and web use cases
- **Maintenance Efficiency**: Bug fixes and improvements automatically benefit all platforms
- **Testing Consistency**: Comprehensive test suite validates behavior across all deployment targets

#### Development Workflow
- **Language Consistency**: Developers work in Go across all components
- **Type Safety**: Compile-time guarantees prevent runtime errors in the web interface
- **Ecosystem Integration**: Leverages Go's extensive standard library and tooling

#### Deployment Benefits
- **No Runtime Dependencies**: Self-contained WASM module requires no external libraries
- **Offline Capability**: Complete functionality available without network connectivity
- **Security**: Sandboxed execution environment with minimal attack surface

### 2.7 Technical Considerations

#### Memory Management
WebAssembly modules have isolated memory spaces that are managed by the Go runtime. Memory allocation and garbage collection operate independently of the browser's JavaScript engine, providing predictable performance characteristics.

#### Error Handling
Error conditions in the WASM module are propagated back to JavaScript through the interop layer, enabling the web interface to provide appropriate user feedback.

#### Future Extensibility
The WASI architecture provides a foundation for adding additional features such as:
- Streaming parsing for large inputs
- Real-time validation and syntax highlighting
- Advanced formatting options
- Custom output formats

---

## 3. Conclusion

LombokToJson demonstrates the effectiveness of purpose-built tokenization combined with innovative cross-platform deployment strategies. The custom tokenizer provides efficient, context-sensitive parsing optimized for Lombok's specific syntax, while the WebAssembly integration enables code reuse across multiple deployment targets without sacrificing performance or functionality.

The architectural decisions reflect a pragmatic approach to software engineering: specialized tools for specific problems, combined with modern deployment technologies that maximize developer productivity and end-user experience. The manual asset management strategy, while seemingly simple, provides the control and reliability necessary for production web applications.

Future development may explore automated build pipelines, streaming compilation for larger inputs, and expanded format support, but the current architecture provides a solid foundation for these enhancements while maintaining the system's core strengths of simplicity, performance, and reliability.

---

## Technical Specifications

- **Language**: Go 1.21+
- **WebAssembly Target**: `js/wasm`
- **Browser Requirements**: WebAssembly 1.0 support
- **Module Size**: ~2-3MB (including Go runtime)
- **Performance**: O(n) parsing complexity
- **Memory Usage**: Linear with input size
- **Supported Platforms**: CLI (cross-platform), Go library, Web browsers

## References

- [Go WebAssembly Documentation](https://golang.org/doc/install/source#environment)
- [WebAssembly Specification](https://webassembly.github.io/spec/)
- [Lombok Project Documentation](https://projectlombok.org/)
- [Monaco Editor Integration](https://microsoft.github.io/monaco-editor/)