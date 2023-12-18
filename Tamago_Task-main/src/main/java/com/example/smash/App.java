package com.example.smash;
import javafx.application.*;
import javafx.geometry.Insets;
import javafx.scene.layout.*;
import javafx.scene.text.Font;
import javafx.scene.*;
import javafx.scene.control.*;
import javafx.stage.*;
import javafx.util.Callback;
import javafx.collections.ObservableList;
import javafx.collections.FXCollections;
import javafx.event.*;
import javafx.scene.control.cell.PropertyValueFactory;
import javafx.scene.control.cell.TextFieldTableCell;
import javafx.scene.image.Image;
import javafx.scene.image.ImageView;

import java.awt.Desktop;
import java.io.File;
import java.io.FileInputStream;
import java.io.IOException;
import java.io.InputStream;
import java.util.logging.*;


/**
 * JavaFX App
 */
public class App extends Application {

    private static Scene scene;
    private static Scene scene2;
    private boolean changeScene = false;
    Stage window;
    private final TableView<Task> tableChron = new TableView<Task>();
    private TableView<Item> tableChron2 = new TableView<Item>();
    private Desktop desktop = Desktop.getDesktop();
    private final ObservableList<Task> data = FXCollections.observableArrayList();
    private int money = 0;
    private ObservableList<Item> data2 = FXCollections.observableArrayList();
    @Override
    public void start(Stage stage) throws IOException {
        window = stage;
        VBox vBox1 = new VBox(); // Create new vertical box.
        VBox leftControl  = new VBox();
        VBox rightControl  = new VBox();
        SplitPane layout2 = new SplitPane();
        Label response = new Label("Menu");
	    MenuBar menuBar = new MenuBar();
        
        Menu fileMenu = new Menu("_File"); // Alt-f shortcut for file 
        MenuItem shop = new MenuItem("Shop");
        MenuItem inventory = new MenuItem("Switch Panel");
        MenuItem exit = new MenuItem("Exit");
        SeparatorMenuItem separator = new SeparatorMenuItem();
               
        fileMenu.getItems().add(shop);
        fileMenu.getItems().add(inventory);
        fileMenu.getItems().add(separator);
        fileMenu.getItems().add(exit);
                        
        menuBar.getMenus().add(fileMenu);

        EventHandler<ActionEvent> MEHandler = new EventHandler<ActionEvent>() {
            public void handle(ActionEvent ae) {
                String name = ((MenuItem)ae.getTarget()).getText();
 
                // if Exit is chosen, the program is  terminated.
                if(name.equals("Exit")) Platform.exit();
 
                response.setText( name + " selected");
            }
        };
        shop.setOnAction(MEHandler);
        inventory.setOnAction(MEHandler);
        exit.setOnAction(MEHandler);
        final FileChooser fileChooser = new FileChooser();
        
        shop.setOnAction(new EventHandler<ActionEvent>(){
                @Override
                public void handle(ActionEvent e) {
                    Shop.display(money);
                    money = Shop.getMoney2();
                    data2.clear();
                    data2.addAll(Shop.getInventory());
                }   
            });
        inventory.setOnAction(new EventHandler<ActionEvent>(){
                @Override
                public void handle(ActionEvent e) {
                    if (!changeScene){
                        changeScene = true;
                        layout2.getItems().addAll(menuBar);
                        window.setScene(scene2);
                    } else if (changeScene){
                        changeScene = false;
                        vBox1.getChildren().addAll(menuBar);
                        window.setScene(scene);
                    }
                    
                }   
            });
                final Label label = new Label("Tamago Task"); // Main window heading label.
                label.setFont(new Font("Arial", 18));
                label.setPadding(new Insets(10, 10, 10, 10));
                Label label2 = new Label("-"); // Main window heading label.
                label2.setFont(new Font("Arial", 18));
                label2.setPadding(new Insets(10, 10, 10, 10));
            
                    tableChron.setEditable(true);
        
                    // Column headings in the tableChron.
                TableColumn<Task, String> taskCol = new TableColumn<Task, String>("Task");
                TableColumn<Task, String> dayCol = new TableColumn<Task, String>("Day");
                TableColumn<Task, String> priceCol = new TableColumn<Task, String>("Price");
                TableColumn<Task, String> deadlineCol = new TableColumn<Task, String>("Deadline");
                TableColumn<Task, String> descriptionCol = new TableColumn<Task, String>("Task Description");
                TableColumn<Task, String> actionCol = new TableColumn<Task, String>("Action");


                TableColumn<Item, String> taskCol2 = new TableColumn<Item, String>("name");
                taskCol2.setCellValueFactory(new PropertyValueFactory<Item, String>("name"));
                taskCol2.setCellFactory(TextFieldTableCell.forTableColumn());
                TableColumn<Item, String> descriptionCol2 = new TableColumn<Item, String>("Description");
                descriptionCol2.setCellValueFactory(new PropertyValueFactory<Item, String>("Description"));
                descriptionCol2.setCellFactory(TextFieldTableCell.forTableColumn());

                taskCol.setCellValueFactory(new PropertyValueFactory<Task, String>("Task"));
                taskCol.setCellFactory(TextFieldTableCell.forTableColumn());
                taskCol.setOnEditCommit(
                    new EventHandler<TableColumn.CellEditEvent<Task, String>>() {
                        @Override
                        public void handle(TableColumn.CellEditEvent<Task, String> t) {
                            ((Task) t.getTableView().getItems().get(
                                t.getTablePosition().getRow())
                                ).setToDo(t.getNewValue());
                        }
                    }
                );
                
                dayCol.setCellValueFactory(new PropertyValueFactory<Task, String>("Day"));
                dayCol.setCellFactory(TextFieldTableCell.forTableColumn());
                dayCol.setOnEditCommit(
                    new EventHandler<TableColumn.CellEditEvent<Task, String>>() {
                        @Override
                        public void handle(TableColumn.CellEditEvent<Task, String> t) {
                            ((Task) t.getTableView().getItems().get(
                                t.getTablePosition().getRow())
                                ).setDay(t.getNewValue());
                        }
                    }
                );

                priceCol.setCellValueFactory(new PropertyValueFactory<Task, String>("Price"));
                priceCol.setCellFactory(TextFieldTableCell.forTableColumn());
                priceCol.setOnEditCommit(
                    new EventHandler<TableColumn.CellEditEvent<Task, String>>() {
                        @Override
                        public void handle(TableColumn.CellEditEvent<Task, String> t) {
                            ((Task) t.getTableView().getItems().get(
                                t.getTablePosition().getRow())
                                ).setPrice(t.getNewValue());
                        }
                    }
                );
                
                deadlineCol.setCellValueFactory(new PropertyValueFactory<Task, String>("Deadline"));
                deadlineCol.setCellFactory(TextFieldTableCell.forTableColumn());
                deadlineCol.setOnEditCommit(
                    new EventHandler<TableColumn.CellEditEvent<Task, String>>() {
                        @Override
                        public void handle(TableColumn.CellEditEvent<Task, String> t) {
                            ((Task) t.getTableView().getItems().get(
                                t.getTablePosition().getRow())
                                ).setDeadline(t.getNewValue());
                        }
                    }
                );
                
                descriptionCol.setCellValueFactory(new PropertyValueFactory<Task, String>("Description"));
                descriptionCol.setCellFactory(TextFieldTableCell.forTableColumn());
                descriptionCol.setOnEditCommit(
                    new EventHandler<TableColumn.CellEditEvent<Task, String>>() {
                        @Override
                        public void handle(TableColumn.CellEditEvent<Task, String> t) {
                            ((Task) t.getTableView().getItems().get(
                                t.getTablePosition().getRow())
                                ).setDescription(t.getNewValue());
                        }
                    }
                );
                Callback<TableColumn<Task, String>, TableCell<Task, String>> cellFactory =
                new Callback<TableColumn<Task, String>, TableCell<Task, String>>() {
            @Override
            public TableCell call(final TableColumn<Task, String> param) {
                final TableCell<Task, String> cell = new TableCell<Task, String>() {

                    final Button btn = new Button("Check");

                    @Override
                    public void updateItem(String item, boolean empty) {
                        super.updateItem(item, empty);
                        if (empty) {
                            setGraphic(null);
                            setText(null);
                        } else {
                            btn.setOnAction(event -> {
                                Task person = getTableView().getItems().get(getIndex());
                                money = money + Integer.parseInt(person.getPrice());
                                System.out.println(money+ "   " + person.getToDo());
                                data.remove(person);
                            });
                            setGraphic(btn);
                            setText(null);
                        }
                    }
                };
                return cell;
            }
        };
        actionCol.setCellFactory(cellFactory);
            tableChron.setItems(data);
            tableChron2.setItems(data2);
            tableChron.getColumns().addAll(taskCol, dayCol,priceCol, deadlineCol, descriptionCol,actionCol); // Place the column headings in tableChron.
            tableChron2.getColumns().addAll(taskCol2, descriptionCol2);
            // Set width
            tableChron.setColumnResizePolicy(TableView.CONSTRAINED_RESIZE_POLICY_FLEX_LAST_COLUMN);
            taskCol.setMinWidth(150);
            dayCol.setMinWidth(70);
            deadlineCol.setMinWidth(60);
            descriptionCol.setMinWidth(200);
            taskCol2.setMinWidth(150);
            descriptionCol2.setMinWidth(200);

        HBox hbox1 = new HBox();
	    hbox1.setSpacing(8);
        hbox1.setPadding(new Insets(10, 10, 10, 10));
        
        // Create text fields so the user can enter new tasks into the table.
	TextField addToDo = new TextField();
	TextField addDay = new TextField();
    TextField addPrice = new TextField();
    TextField addDeadline = new TextField();
	TextField addDescription = new TextField();
    addToDo.setPromptText("Enter Task name");
    addDay.setPromptText("Task date");
    addPrice.setPromptText("Task price");
    addDeadline.setPromptText("Task deadline");
    addDescription.setPromptText("Task description");

    // Set length
        addToDo.setPrefWidth(200);
        addDay.setPrefWidth(150);
        addPrice.setPrefWidth(150);
        addDeadline.setPrefWidth(150);
        addDescription.setPrefWidth(200);
        
	// Create add button. 
	Button btnAdd = new Button("Add");
    btnAdd.setMinWidth(50);
    Button remove = new Button("Remove");
    remove.setMinWidth(80);
    Button set = new Button("Set");
    btnAdd.setMinWidth(50);
	btnAdd.setOnAction(new EventHandler<ActionEvent>() {
            @Override
            public void handle(ActionEvent e) {
                Task task = new Task(addToDo.getText(), addDay.getText(),Integer.parseInt(addPrice.getText()), addDeadline.getText(), addDescription.getText());
                data.add(task);
                addToDo.clear();
                addDay.clear();
                addPrice.clear();
                addDeadline.clear();
                addDescription.clear();
            }
        });
    remove.setOnAction(e -> {
        int person = tableChron.getSelectionModel().getSelectedIndex();
        if (person != -1){
            data.remove(person);
        }
    });
    set.setOnAction(e -> {
        Item person = tableChron2.getSelectionModel().getSelectedItem();
        if (person != null){
            label2.setText(person.getToDo());
        }
    });
    InputStream stream = new FileInputStream("./src/main/resources/com/example/smash/sonic1.gif");
    Image image = new Image(stream);
    //Creating the image view
    ImageView imageView = new ImageView();
    //Setting image to the image view
    imageView.setImage(image);
    //Setting the image view parameters
    imageView.setX(10);
    imageView.setY(10);
    imageView.setFitWidth(200);
    imageView.setPreserveRatio(true);
    // Get user entry fields all in a horizontal row.
	hbox1.getChildren().addAll(addToDo, addDay,addPrice, addDeadline, addDescription, btnAdd,remove);
    leftControl.getChildren().addAll(imageView,label2);
    rightControl.getChildren().addAll(tableChron2);
	// Get all items in a vertical, stacking format.
	vBox1.getChildren().addAll(menuBar, label, tableChron, hbox1);
    layout2.getItems().addAll(menuBar,leftControl,rightControl,set);
        scene = new Scene(vBox1, 640, 480);
        scene2 = new Scene(layout2, 640, 480);
        window.setTitle("Tamago Task");
        window.setScene(scene);
        window.show();
    }

    public static void main(String[] args) {
        launch();
    }

}