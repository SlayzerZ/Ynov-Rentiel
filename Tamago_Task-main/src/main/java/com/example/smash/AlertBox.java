package com.example.smash;

import javafx.geometry.Pos;
import javafx.scene.Scene;
import javafx.scene.control.Button;
import javafx.scene.control.Label;
import javafx.scene.layout.VBox;
import javafx.stage.Modality;
import javafx.stage.Stage;

public class AlertBox {
    public static void display(String title, String message){
        Stage window = new Stage();
        window.initModality(Modality.APPLICATION_MODAL);
        window.setTitle(title);
        window.setMinWidth(250);
        Label label = new Label();
        label.setText(message);
        Button closeButton = new Button("fermer");
        closeButton.setOnAction(e -> window.close());
        VBox layout1 = new VBox(10);
        layout1.getChildren().addAll(label, closeButton);
        layout1.setAlignment(Pos.CENTER);
        Scene scene = new Scene(layout1);
        window.setScene(scene);
        window.showAndWait();
    }
}
