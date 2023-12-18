package com.example.smash;



import java.util.ArrayList;

import javafx.application.Platform;
import javafx.collections.FXCollections;
import javafx.collections.ObservableList;
import javafx.geometry.Insets;
import javafx.scene.Scene;
import javafx.scene.control.Button;
import javafx.scene.control.Label;
import javafx.scene.control.TableCell;
import javafx.scene.control.TableColumn;
import javafx.scene.control.TableView;
import javafx.scene.control.cell.PropertyValueFactory;
import javafx.scene.control.cell.TextFieldTableCell;
import javafx.scene.layout.VBox;
import javafx.scene.text.Font;
import javafx.stage.Modality;
import javafx.stage.Stage;
import javafx.util.Callback;
//import prerna.sablecc2.reactor.app.metaeditor.ReloadAppOwlReactor;

public class Shop {
    static Label label = new Label(); // Main window heading label.
    private static Item item1 = new Item("Una Merda Alta", 150, "T'es une grosse merde") ;
    private static Item item2 = new Item("L is real", 1500, "L est réel") ;
    private static Item item3 = new Item("Two becomes One", 3000, "Deux devient un") ;
    private static Item item4 = new Item("Nanotech", 4500, "T'as le stock max de nanotech") ;
    private static Item item5 = new Item("The World thats never was", 10000, "Le monde qui n'a jamais existé") ;
    private static Item item6 = new Item("Sixth Badge", 20000, "Tu as obtenu le badge je me fais chier") ;
    private static Item item7 = new Item("EXPLOSION", 30000, "EXPLOSION!!!") ;
    private static Item item8 = new Item("Dragon balls", 40000, "Tu peux réaliser un souhait") ;
    private static Item item9 = new Item("The Chosen One", 50000, "Tu es l'élu") ;
    public static int money2;
    public static ArrayList<Item> inventory = new ArrayList<Item>();
    public static void setLabel(String lab){
        label.setText(lab);
    }
    public void setMoney2(int money2){
        money2 = this.money2;
    }
    public static int getMoney2(){
        return money2;
    }
    public static ArrayList<Item> getInventory(){
        return inventory;
    }
    public static void display(int money){
        money2 = money;
        final TableView<Item> tableChron = new TableView<Item>();
        final ObservableList<Item> data = FXCollections.observableArrayList();
        Stage window = new Stage();
        window.initModality(Modality.APPLICATION_MODAL);
        window.setTitle("Shop");
        window.setMinWidth(250);
        label.setText("Your money is " + Integer.toString(money2) + " Bolts");
        label.setFont(new Font("Arial", 18));
        label.setPadding(new Insets(10, 10, 10, 10));
        VBox layout1 = new VBox();
        Scene scene = new Scene(layout1);
        Button closeButton = new Button("Fermer");
        TableColumn<Item, String> taskCol = new TableColumn<Item, String>("name");
        taskCol.setCellValueFactory(new PropertyValueFactory<Item, String>("name"));
        taskCol.setCellFactory(TextFieldTableCell.forTableColumn());
        TableColumn<Item, String> priceCol = new TableColumn<Item, String>("Price");
        priceCol.setCellValueFactory(new PropertyValueFactory<Item, String>("Price"));
        priceCol.setCellFactory(TextFieldTableCell.forTableColumn());
        TableColumn<Item, String> descriptionCol = new TableColumn<Item, String>("Description");
        descriptionCol.setCellValueFactory(new PropertyValueFactory<Item, String>("Description"));
        descriptionCol.setCellFactory(TextFieldTableCell.forTableColumn());
        TableColumn<Item, String> buyCol = new TableColumn<Item, String>("Buy");
        Callback<TableColumn<Item, String>, TableCell<Item, String>> cellFactory =
                new Callback<TableColumn<Item, String>, TableCell<Item, String>>() {
            @Override
            public TableCell call(final TableColumn<Item, String> param) {
                final TableCell<Item, String> cell = new TableCell<Item, String>() {

                    final Button btn = new Button("BUY!");

                    @Override
                    public void updateItem(String item, boolean empty) {
                        super.updateItem(item, empty);
                        if (empty) {
                            setGraphic(null);
                            setText(null);
                        } else {
                            btn.setOnAction(event -> {
                                Item person = getTableView().getItems().get(getIndex());
                                if (money2 >= Integer.parseInt(person.getPrice())){
                                    money2 = money2 - Integer.parseInt(person.getPrice());
                                    inventory.add(person);
                                    setLabel("Your money is " + Integer.toString(money2) + " Bolts");
                                } else {
                                    AlertBox.display("Not Enough", "Your actual Wallet is not enough");
                                }
                            });
                            setGraphic(btn);
                            setText(null);
                        }
                    }
                };
                return cell;
            }
        };
        buyCol.setCellFactory(cellFactory);
        data.addAll(item1,item2,item3,item4,item5,item6,item7,item8,item9);
        tableChron.setItems(data);
        tableChron.getColumns().addAll(taskCol, priceCol, descriptionCol,buyCol);
        taskCol.setMinWidth(150);
        descriptionCol.setMinWidth(200);
        closeButton.setOnAction(e ->{
           // layout1.getChildren().removeAll(tableChron);
            window.close();
        });
        layout1.getChildren().addAll(label,closeButton,tableChron);
        window.setScene(scene);
        window.showAndWait();
    }
}
