# tc-create-app.375
Text of [issue](https://github.com/unfoldingWord/tc-create-app/issues/375):

```
v1.0.3  
There is an inconsistency in providing an "X" to close popup menus  
The View Columns and Filter Table menus show an X to allow the user to close them. The Manage Versions menu does not.  
Suggest adding the X to the Manage Versions menu.

The Manage versions dropdown should have an X like the other dropdowns
```


Found the code for the popup here in 
`.\scripture-resources-rcl\src\components\parallel-scripture\ColumnsMenu.js`
```js
    <Menu anchorEl={anchorEl} open={!!anchorEl} onClose={() => onAnchorEl()}>
      <MenuItem
        key={'text'}
        disabled
        style={{
          opacity: 1, fontWeight: 600, fontSize: 12,
        }}
      >
        {localString('ManageVersions')}
      </MenuItem>

      {menuItems}

      <MenuItem>
        <TextField
          id='resourceUrl'
          label={localString('ResourcePath')}
          variant='outlined'
          defaultValue=''
          style={{ valign: 'middle' }}
        />
        <Tooltip title={localString('AddResource')} arrow>
          <IconButton
            aria-label={localString('AddResource')}
            onClick={onResourceAddClick}
            className={(classes.action, classes.menuIconButton)}
            size='small'
          >
            <PlaylistAdd fontSize='small' />
          </IconButton>
        </Tooltip>
      </MenuItem>
    </Menu>
```

The `Menu` component is a standard '@material-ui/core' component...

Based on the material ui documentation, there is no option to add a menu close button. Could try a button at the bottom to "close" it. The onClick action would be to run the same onClose action on the Menu element.

1. run script `set-branch feature-cn-375-add-close-to-addResource-menu`
2. researched a couple of ways of doing this.
	- adding a close button as last "menu item"
	- adding an "X" (as the issue directs)
3. found that button works ok. It's onClick function is simply the same as the Menu's onClose function.
4. found that material UI does not support a "close" feature. But eventually we were able to simulate it by adding a menu item before the title that contained a right justified icon button.

Here is the code:
```js
      <MenuItem style={{justifyContent: "flex-end"}}>
        <IconButton 
          aria-label='Close' 
          onClick={() => {onAnchorEl()}} 
          className={classes.close}
          disabled={false}
        >
          <CloseIcon fontSize='small' /> 
        </IconButton>
      </MenuItem>
```

The className "close" was added at the end to remove padding:
```js
const useStyles = makeStyles((theme) => ({
  root: {},
  action: { padding: '8px' },
  close: { padding: '0px' },
}));
```