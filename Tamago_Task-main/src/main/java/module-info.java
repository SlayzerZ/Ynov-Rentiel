module com.example.smash {
    requires javafx.controls;
    requires javafx.fxml;
    requires transitive javafx.graphics;
    requires java.logging;
    requires java.desktop;
    opens com.example.smash to javafx.fxml;
    exports com.example.smash;
}
